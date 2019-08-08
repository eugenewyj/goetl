// Copyright 2018 The GoEtl Authors. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package etl

// Step 定义了一个基本ETL（抽取、转换、加载）操作包括的方法。
type Step interface {
	Execute() (err error)
}
