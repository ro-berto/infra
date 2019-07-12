// Copyright 2019 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clients

import (
	"fmt"
	"net/url"

	"go.chromium.org/gae/service/taskqueue"
	"go.chromium.org/luci/common/logging"
	"golang.org/x/net/context"
)

const repairBotsQueue = "repair-bots"
const resetBotsQueue = "reset-bots"

// PushRepairDUTs pushes duts for taskqueue repair-bots upcoming repair jobs.
func PushRepairDUTs(ctx context.Context, dutNames []string) error {
	return pushDUTs(ctx, dutNames, repairBotsQueue, repairTask)
}

// PushResetDUTs pushes duts for taskqueue reset-bots upcoming reset jobs.
func PushResetDUTs(ctx context.Context, dutNames []string) error {
	return pushDUTs(ctx, dutNames, resetBotsQueue, resetTask)
}

func repairTask(dn string) *taskqueue.Task {
	values := url.Values{}
	values.Set("dutName", dn)
	return taskqueue.NewPOSTTask(fmt.Sprintf("/internal/task/repair/%s", dn), values)
}

func resetTask(dn string) *taskqueue.Task {
	values := url.Values{}
	values.Set("dutName", dn)
	return taskqueue.NewPOSTTask(fmt.Sprintf("/internal/task/reset/%s", dn), values)
}

func pushDUTs(ctx context.Context, dutNames []string, queueName string, taskGenerator func(string) *taskqueue.Task) error {
	tasks := make([]*taskqueue.Task, 0, len(dutNames))
	for _, dn := range dutNames {
		tasks = append(tasks, taskGenerator(dn))
	}
	if err := taskqueue.Add(ctx, queueName, tasks...); err != nil {
		return err
	}
	logging.Infof(ctx, "enqueued %d tasks", len(tasks))
	return nil
}
