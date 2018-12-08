// Copyright 2018 The goetl Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main 是goetl应用入口(github.com/eugenewyj/goetl/goetlmain)的
// 一个简单封装，保证goetl能够通过`go get github.com/eugenewyj/goetl`获取，
// 并在$GOBIN/goetl下生成二进制文件。
//
// 这个包不允许以任何形式进行修改或扩展；为了修改goetl二进制程序，请修改
// github.com/eugenewyj/goetl/goetlmain中的内容。
//
package main

import "github.com/eugenewyj/goetl/goetlmain"

func main() {
	goetlmain.Main()
}
