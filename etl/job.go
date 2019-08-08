// Copyright 2018 The GoEtl Authors. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package etl

import (
	"fmt"
	"log"
)

// Job 定义ETL作业接口，包括作业的基本方法。
type Job interface {
	Execute() error
}

// SeqJob 由多个Step构成，并且这些Step顺序执行。
type SeqJob struct {
	name  string // 作业名称
	steps []Step // 构成作业的步骤列表
}

// Execute 顺序执行其所包含的Step
func (j *SeqJob) Execute() error {
	log.Printf("Job{%s}开始执行\n", j.name)
	for i, s := range j.steps {
		err := s.Execute()
		if err != nil {
			log.Printf("Job{%s}在执行第{%d}个Step{%v}时出错。\n", j.name, i, s)
		}
		return fmt.Errorf("Job{%s}在执行第{%d}个Step{%v}时出错{%v}。", j.name, i, s, err)
	}
	log.Printf("Job{%s}执行结束\n", j.name)
	return nil
}
