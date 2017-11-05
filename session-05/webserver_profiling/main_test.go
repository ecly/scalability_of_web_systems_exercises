// Copyright 2017 Google Inc.
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

package main

import "testing"
import "strings"

var name string
var ok bool

func BenchmarkIsGopher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name, ok = isGopher("campoy@golang.org")
	}
}

//var res []string

func BenchmarkFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//go:noinline
		strings.Fields("hello, dear friend")
	}
}