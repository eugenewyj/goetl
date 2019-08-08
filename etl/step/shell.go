// Copyright 2018 The GoEtl Authors. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package step

// Shell 一个执行Shell脚本的步骤。
type Shell struct {
	path string	// 脚本路径
}

// 执行shell脚本。
func (s *Shell) Execute()  error {
	return nil
}