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

package sum

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	ints := []int{1, 2, 3}
	if res := All(ints...); res != 6 {
		t.Errorf("Sum of %v should be 6 but was %v", ints, res)
	}
}

func TestAll_Subtest(t *testing.T) {
	tt := []struct {
		name string
		vals []int
		res  int
	}{
		{"postive consecutive", []int{1, 2, 3}, 6},
		{"positive and negative", []int{-3, 1, 3}, 1},
		{"negative consecutive", []int{-1, -2, -3}, -6},
		{"zeros", []int{0, 0, 0}, 0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if res := All(tc.vals...); res != tc.res {
				t.Fatalf("Test %s expected %d, but got %d", tc.name, tc.res, res)
			}
		})
	}
}

var s int

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = All(i)
	}
}

func ExampleAll() {
	fmt.Println(All([]int{1, 1, 1}...))
	// Output:
	// 3
}
