// Copyright 2018 The GoEtl Authors. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package etl

import "testing"

// TestJobExecute 测试任务的执行
func TestJobExecute(t *testing.T) {
	var j = &SeqJob{}
	j.Execute()
}