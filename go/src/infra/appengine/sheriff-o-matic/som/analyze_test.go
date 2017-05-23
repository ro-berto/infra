package som

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/net/context"

	"infra/monitoring/client"
	testhelper "infra/monitoring/client/test"
	"infra/monitoring/messages"

	"github.com/luci/gae/impl/dummy"
	"github.com/luci/gae/service/datastore"
	"github.com/luci/gae/service/info"
	"github.com/luci/gae/service/urlfetch"
	"github.com/luci/luci-go/appengine/gaetesting"
	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/server/router"

	. "github.com/smartystreets/goconvey/convey"
)

func newTestContext() context.Context {
	ctx := gaetesting.TestingContext()
	ta := datastore.GetTestable(ctx)
	ta.Consistent(true)
	return ctx
}

type giMock struct {
	info.RawInterface
	token  string
	expiry time.Time
	err    error
}

func (gi giMock) AccessToken(scopes ...string) (token string, expiry time.Time, err error) {
	return gi.token, gi.expiry, gi.err
}

func setUpGitiles(c context.Context) context.Context {
	return urlfetch.Set(c, &testhelper.MockGitilesTransport{
		Responses: map[string]string{
			gkTreesURL: `{    "chromium": {
        "build-db": "waterfall_build_db.json",
        "masters": {
            "https://build.chromium.org/p/chromium": ["*"]
        },
        "open-tree": true,
        "password-file": "/creds/gatekeeper/chromium_status_password",
        "revision-properties": "got_revision_cp",
        "set-status": true,
        "status-url": "https://chromium-status.appspot.com",
        "track-revisions": true
    }}`,
			gkTreesInternalURL: `{    "chromium": {
        "build-db": "waterfall_build_db.json",
        "masters": {
            "https://build.chromium.org/p/chromium": ["*"]
        },
        "open-tree": true,
        "password-file": "/creds/gatekeeper/chromium_status_password",
        "revision-properties": "got_revision_cp",
        "set-status": true,
        "status-url": "https://chromium-status.appspot.com",
        "track-revisions": true
    }}`,
			gkTreesCorpURL: `{    "chromium": {
        "build-db": "waterfall_build_db.json",
        "masters": {
            "https://build.chromium.org/p/chromium": ["*"]
        },
        "open-tree": true,
        "password-file": "/creds/gatekeeper/chromium_status_password",
        "revision-properties": "got_revision_cp",
        "set-status": true,
        "status-url": "https://chromium-status.appspot.com",
        "track-revisions": true
    }}`,
			gkConfigInternalURL: `
{
  "comment": ["This is a configuration file for gatekeeper_ng.py",
              "Look at that for documentation on this file's format."],
  "masters": {
    "https://build.chromium.org/p/chromium": [
      {
        "categories": [
          "chromium_tree_closer"
        ],
        "builders": {
          "Win": {
            "categories": [
              "chromium_windows"
            ]
          },
          "*": {}
        }
      }
    ]
   }
}`,

			gkConfigURL: `
{
  "comment": ["This is a configuration file for gatekeeper_ng.py",
              "Look at that for documentation on this file's format."],
  "masters": {
    "https://build.chromium.org/p/chromium": [
      {
        "categories": [
          "chromium_tree_closer"
        ],
        "builders": {
          "Win": {
            "categories": [
              "chromium_windows"
            ]
          },
          "*": {}
        }
      }
    ]
   }
}`,
			gkConfigCorpURL: `
{
  "comment": ["This is a configuration file for gatekeeper_ng.py",
              "Look at that for documentation on this file's format."],
  "masters": {
    "https://build.chromium.org/p/chromium": [
      {
        "categories": [
          "chromium_tree_closer"
        ],
        "builders": {
          "Win": {
            "categories": [
              "chromium_windows"
            ]
          },
          "*": {}
        }
      }
    ]
   }
}`,
		}})
}

func TestGetAnalyzeHandler(t *testing.T) {
	Convey("bad request", t, func() {
		c := gaetesting.TestingContext()
		c = info.SetFactory(c, func(ic context.Context) info.RawInterface {
			return giMock{dummy.Info(), "", time.Now(), nil}
		})
		c = setUpGitiles(c)
		w := httptest.NewRecorder()

		ctx := &router.Context{
			Context: c,
			Writer:  w,
			Request: makeGetRequest(),
			Params:  makeParams("tree", "unknown.tree"),
		}
		getAnalyzeHandler(ctx)

		So(w.Code, ShouldEqual, http.StatusNotFound)
	})

	Convey("ok request", t, func() {
		c := newTestContext()
		c = info.SetFactory(c, func(ic context.Context) info.RawInterface {
			return giMock{dummy.Info(), "", time.Now(), nil}
		})
		c = setUpGitiles(c)

		c = client.WithReader(c, testhelper.MockReader{
			BuildExtracts: map[string]*messages.BuildExtract{
				"chromium": {},
			},
		})

		w := httptest.NewRecorder()

		ctx := &router.Context{
			Context: c,
			Writer:  w,
			Request: makeGetRequest(),
			Params:  makeParams("tree", "chromium"),
		}
		getAnalyzeHandler(ctx)

		So(w.Code, ShouldEqual, http.StatusOK)
	})

	Convey("ok request, no gitiles", t, func() {
		c := newTestContext()
		c = info.SetFactory(c, func(ic context.Context) info.RawInterface {
			return giMock{dummy.Info(), "", time.Now(), nil}
		})
		c = urlfetch.Set(c, &testhelper.MockGitilesTransport{})

		c = client.WithReader(c, testhelper.MockReader{
			BuildExtracts: map[string]*messages.BuildExtract{
				"chromium": {},
			},
		})

		w := httptest.NewRecorder()

		ctx := &router.Context{
			Context: c,
			Writer:  w,
			Request: makeGetRequest(),
			Params:  makeParams("tree", "chromium"),
		}
		getAnalyzeHandler(ctx)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestStoreAlertsSummary(t *testing.T) {
	Convey("success", t, func() {
		c := gaetesting.TestingContext()
		c = info.SetFactory(c, func(ic context.Context) info.RawInterface {
			return giMock{dummy.Info(), "", clock.Now(c), nil}
		})
		c = setUpGitiles(c)
		err := storeAlertsSummary(c, nil, "some tree", &messages.AlertsSummary{
			Alerts: []messages.Alert{
				{
					Title:     "foo",
					Extension: messages.BuildFailure{},
				},
			},
		})
		So(err, ShouldBeNil)
	})
}

func TestGetMiloDiffHandler(t *testing.T) {
	Convey("MiloDiffHandler", t, func() {
		c := gaetesting.TestingContext()
		c = info.SetFactory(c, func(ic context.Context) info.RawInterface {
			return giMock{dummy.Info(), "", time.Now(), nil}
		})
		c = setUpGitiles(c)
		w := httptest.NewRecorder()

		Convey("bad request", func() {
			ctx := &router.Context{
				Context: c,
				Writer:  w,
				Request: makeGetRequest(),
				Params:  makeParams("tree", "unknown.tree"),
			}
			getMiloDiffHandler(ctx)

			So(w.Code, ShouldEqual, http.StatusNotFound)
		})

		Convey("ok request", func() {
			c = client.WithReader(c, testhelper.MockReader{
				BuildExtracts: map[string]*messages.BuildExtract{
					"chromium": {},
				},
			})

			ta := datastore.GetTestable(c)
			ta.AddIndexes(&datastore.IndexDefinition{
				Kind:     "AlertsJSON",
				Ancestor: true,
				SortBy: []datastore.IndexColumn{
					{
						Property:   "Date",
						Descending: true,
					},
				},
			})

			alertsJSON := &AlertsJSON{
				Tree:     datastore.MakeKey(c, "Tree", "chromium"),
				Date:     clock.Now(c).UTC(),
				Contents: []byte("{}"),
			}

			err := datastore.Put(c, alertsJSON)
			So(err, ShouldBeNil)

			alertsJSONMilo := &AlertsJSON{
				Tree:     datastore.MakeKey(c, "Tree", "milo.chromium"),
				Date:     clock.Now(c).UTC(),
				Contents: []byte("{'some': 'value'}"),
			}

			err = datastore.Put(c, alertsJSONMilo)
			So(err, ShouldBeNil)

			ctx := &router.Context{
				Context: c,
				Writer:  w,
				Request: makeGetRequest(),
				Params:  makeParams("tree", "chromium"),
			}

			alertsJSON, err = getAlertsForTree(c, "chromium")
			So(err, ShouldBeNil)

			getMiloDiffHandler(ctx)
			So(w.Code, ShouldEqual, http.StatusOK)
		})

		Convey("ok request, no gitiles", func() {
			c = info.SetFactory(c, func(ic context.Context) info.RawInterface {
				return giMock{dummy.Info(), "", time.Now(), nil}
			})
			c = urlfetch.Set(c, &testhelper.MockGitilesTransport{})

			c = client.WithReader(c, testhelper.MockReader{
				BuildExtracts: map[string]*messages.BuildExtract{
					"chromium": {},
				},
			})

			ctx := &router.Context{
				Context: c,
				Writer:  w,
				Request: makeGetRequest(),
				Params:  makeParams("tree", "chromium"),
			}
			getMiloDiffHandler(ctx)

			So(w.Code, ShouldEqual, http.StatusInternalServerError)
		})
	})
}
