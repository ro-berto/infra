// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crauditcommits

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/context"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/gae/impl/memory"
	ds "go.chromium.org/gae/service/datastore"
	"go.chromium.org/gae/service/mail"
	"go.chromium.org/gae/service/user"
	"go.chromium.org/luci/server/router"

	"infra/monorail"
)

// sendEmailForFinditViolation is not actually used by any RuleSet its purpose
// is to illustrate how one would use sendEmailForViolation to notify about
// violations via email.
func sendEmailForFinditViolation(ctx context.Context, cfg *RepoConfig, rc *RelevantCommit, cs *Clients, state string) (string, error) {
	recipients := []string{"eng-team@dummy.com"}
	subject := "A policy violation was detected on commit %s"
	return sendEmailForViolation(ctx, cfg, rc, cs, state, recipients, subject)
}

func TestNotifier(t *testing.T) {

	Convey("ViolationNotifier handler test", t, func() {
		ctx := memory.UseWithAppID(context.Background(), "cr-audit-commits-test")

		notifierPath := "/_cron/violationnotifier"

		withTestingContext := func(c *router.Context, next router.Handler) {
			c.Context = ctx
			ds.GetTestable(ctx).CatchupIndexes()
			next(c)
		}
		user.GetTestable(ctx).Login("notifier@cr-audit-commits-test.appspotmail.com", "", false)

		r := router.New()
		r.GET(notifierPath, router.NewMiddlewareChain(withTestingContext), ViolationNotifier)
		srv := httptest.NewServer(r)
		client := &http.Client{}
		testClients = &Clients{}
		Convey("Unknown Repo", func() {
			resp, err := client.Get(srv.URL + notifierPath + "?repo=unknown")
			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, 500)

		})
		Convey("New Repo", func() {
			RuleMap["new-repo"] = &RepoConfig{
				BaseRepoURL:     "https://new.googlesource.com/new.git",
				GerritURL:       "https://new-review.googlesource.com",
				BranchName:      "master",
				StartingCommit:  "000000",
				MonorailAPIURL:  "https://monorail-fake.appspot.com/_ah/api/monorail/v1",
				MonorailProject: "fakeproject",
				Rules: map[string]RuleSet{"rules": AccountRules{
					Account: "new@test.com",
					Funcs: []RuleFunc{func(c context.Context, ap *AuditParams, rc *RelevantCommit, cs *Clients) *RuleResult {
						return &RuleResult{"Dummy rule", rulePassed, ""}
					}},
					notificationFunction: fileBugForFinditViolation,
				}},
			}
			Convey("Should fail", func() {
				resp, err := client.Get(srv.URL + notifierPath + "?repo=new-repo")
				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldNotEqual, 200)
			})
		})
		Convey("Existing Repo", func() {
			RuleMap["old-repo"] = &RepoConfig{
				BaseRepoURL:     "https://old.googlesource.com/old.git",
				GerritURL:       "https://old-review.googlesource.com",
				BranchName:      "master",
				StartingCommit:  "000000",
				MonorailAPIURL:  "https://monorail-fake.appspot.com/_ah/api/monorail/v1",
				MonorailProject: "fakeproject",
				NotifierEmail:   "notifier@cr-audit-commits-test.appspotmail.com",
				Rules: map[string]RuleSet{"rules": AccountRules{
					Account: "author@test.com",
					Funcs: []RuleFunc{func(c context.Context, ap *AuditParams, rc *RelevantCommit, cs *Clients) *RuleResult {
						return &RuleResult{"Dummy rule", rulePassed, ""}
					}},
					notificationFunction: fileBugForFinditViolation,
				}},
			}
			repoState := &RepoState{
				RepoURL:            "https://old.googlesource.com/old.git/+/master",
				LastKnownCommit:    "123456",
				LastRelevantCommit: "999999",
			}
			ds.Put(ctx, repoState)

			Convey("No audits", func() {
				testClients.monorail = mockMonorailClient{
					e: fmt.Errorf("Monorail was called even though there were no failed audits"),
				}
				resp, err := client.Get(srv.URL + notifierPath + "?repo=old-repo")
				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, 200)
			})
			Convey("No failed audits", func() {
				rsk := ds.KeyForObj(ctx, repoState)
				testClients.monorail = mockMonorailClient{
					e: fmt.Errorf("Monorail was called even though there were no failed audits"),
				}
				rc := &RelevantCommit{
					RepoStateKey:     rsk,
					CommitHash:       "600dc0de",
					Status:           auditCompleted,
					Result:           []RuleResult{{"DummyRule", rulePassed, ""}},
					CommitterAccount: "committer@test.com",
					AuthorAccount:    "author@test.com",
					CommitMessage:    "This commit passed all audits.",
				}
				err := ds.Put(ctx, rc)
				So(err, ShouldBeNil)

				resp, err := client.Get(srv.URL + notifierPath + "?repo=old-repo")
				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, 200)
				rc = &RelevantCommit{
					RepoStateKey: rsk,
					CommitHash:   "600dc0de",
				}
				err = ds.Get(ctx, rc)
				So(err, ShouldBeNil)
				So(rc.GetNotificationState("rules"), ShouldEqual, "")
				So(rc.NotifiedAll, ShouldBeFalse)
			})
			Convey("Failed audits - bug only", func() {
				rsk := ds.KeyForObj(ctx, repoState)
				testClients.monorail = mockMonorailClient{
					il: &monorail.IssuesListResponse{},
					ii: &monorail.InsertIssueResponse{
						Issue: &monorail.Issue{
							Id: 12345,
						},
					},
				}
				rc := &RelevantCommit{
					RepoStateKey:     rsk,
					CommitHash:       "badc0de",
					Status:           auditCompletedWithViolation,
					Result:           []RuleResult{{"DummyRule", ruleFailed, "This commit is bad"}},
					CommitterAccount: "committer@test.com",
					AuthorAccount:    "author@test.com",
					CommitMessage:    "This commit failed all audits.",
				}
				err := ds.Put(ctx, rc)
				So(err, ShouldBeNil)

				resp, err := client.Get(srv.URL + notifierPath + "?repo=old-repo")
				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, 200)
				rc = &RelevantCommit{
					RepoStateKey: rsk,
					CommitHash:   "badc0de",
				}
				err = ds.Get(ctx, rc)
				So(err, ShouldBeNil)
				So(rc.GetNotificationState("rules"), ShouldEqual, "BUG=12345")
				So(rc.NotifiedAll, ShouldBeTrue)
				m := mail.GetTestable(ctx)
				So(m.SentMessages(), ShouldBeEmpty)

			})
			Convey("Exceeded retries", func() {
				rsk := ds.KeyForObj(ctx, repoState)
				testClients.monorail = mockMonorailClient{
					ii: &monorail.InsertIssueResponse{
						Issue: &monorail.Issue{
							Id: 12345,
						},
					},
				}
				rc := &RelevantCommit{
					RepoStateKey:     rsk,
					CommitHash:       "b00b00",
					Status:           auditFailed,
					Result:           []RuleResult{},
					CommitterAccount: "committer@test.com",
					AuthorAccount:    "author@test.com",
					CommitMessage:    "This commit panicked and panicked",
					Retries:          MaxRetriesPerCommit + 1,
				}
				err := ds.Put(ctx, rc)
				So(err, ShouldBeNil)

				resp, err := client.Get(srv.URL + notifierPath + "?repo=old-repo")
				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, 200)
				rc = &RelevantCommit{
					RepoStateKey: rsk,
					CommitHash:   "b00b00",
				}
				err = ds.Get(ctx, rc)
				So(err, ShouldBeNil)
				So(rc.GetNotificationState("AuditFailure"), ShouldEqual, "BUG=12345")
				So(rc.NotifiedAll, ShouldBeTrue)
			})
		})
		Convey("Failed audits - email only", func() {
			RuleMap["old-repo-email"] = &RepoConfig{
				BaseRepoURL:     "https://old.googlesource.com/old-email.git",
				GerritURL:       "https://old-review.googlesource.com",
				BranchName:      "master",
				StartingCommit:  "000000",
				MonorailAPIURL:  "https://monorail-fake.appspot.com/_ah/api/monorail/v1",
				MonorailProject: "fakeproject",
				NotifierEmail:   "notifier@cr-audit-commits-test.appspotmail.com",
				Rules: map[string]RuleSet{"rulesEmail": AccountRules{
					Account: "author@test.com",
					Funcs: []RuleFunc{func(c context.Context, ap *AuditParams, rc *RelevantCommit, cs *Clients) *RuleResult {
						return &RuleResult{"Dummy rule", rulePassed, ""}
					}},
					notificationFunction: sendEmailForFinditViolation,
				}},
			}
			repoState := &RepoState{
				RepoURL:            "https://old.googlesource.com/old-email.git/+/master",
				LastKnownCommit:    "123456",
				LastRelevantCommit: "999999",
			}
			ds.Put(ctx, repoState)
			rsk := ds.KeyForObj(ctx, repoState)
			rc := &RelevantCommit{
				RepoStateKey:     rsk,
				CommitHash:       "badc0de",
				Status:           auditCompletedWithViolation,
				Result:           []RuleResult{{"DummyRule", ruleFailed, "This commit is bad"}},
				CommitterAccount: "committer@test.com",
				AuthorAccount:    "author@test.com",
				CommitMessage:    "This commit failed all audits.",
			}
			err := ds.Put(ctx, rc)
			So(err, ShouldBeNil)

			resp, err := client.Get(srv.URL + notifierPath + "?repo=old-repo-email")
			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, 200)
			rc = &RelevantCommit{
				RepoStateKey: rsk,
				CommitHash:   "badc0de",
			}
			err = ds.Get(ctx, rc)
			So(err, ShouldBeNil)
			m := mail.GetTestable(ctx)
			So(rc.NotifiedAll, ShouldBeTrue)
			So(m.SentMessages()[0], ShouldResemble,
				&mail.TestMessage{
					Message: mail.Message{
						Sender:  "notifier@cr-audit-commits-test.appspotmail.com",
						To:      []string{"eng-team@dummy.com"},
						Subject: "A policy violation was detected on commit badc0de",
						Body:    "Here are the messages from the rules that were broken by this commit:\n\nThis commit is bad",
					}})

		})
	})
}
