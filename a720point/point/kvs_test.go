// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package point

import (
	"fmt"
	"testing"
)

// $ GODEBUG=gctrace=1 go test -benchmem -bench Benchmark_Add1 -cpu=1
// gc 96 @1.060s 1%: 0.016+2.5+0.002 ms clock, 0.016+0.059/0.077/0+0.002 ms cpu, 3->3->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 1 P
// 10000            105114 ns/op           29609 B/op        979 allocs/op
// gc 180 @1.243s 2%: 0.011+2.6+0.001 ms clock, 0.011+0.12/0/0+0.001 ms cpu, 3->3->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 1 P
// 1000000              1233 ns/op             552 B/op         17 allocs/op
func Benchmark_Add1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testFunc_1()
	}
}

// $ GODEBUG=gctrace=1 go test -benchmem -bench Benchmark_Add2 -cpu=1
// gc 94 @1.088s 1%: 0.012+4.0+0.002 ms clock, 0.012+0.18/0/0+0.002 ms cpu, 3->3->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 1 P
// 10000            108425 ns/op           29609 B/op        979 allocs/op
// gc 180 @1.316s 2%: 0.008+0.24+0 ms clock, 0.008+0.055/0.069/0+0 ms cpu, 3->3->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 1 P
// 1000000              1310 ns/op             552 B/op         17 allocs/op
func Benchmark_Add2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testFunc_2()
	}
}

var lenKV = 2

func testFunc_1() {
	var kvs KVs
	for i := 0; i < lenKV; i++ {
		kvs = kvs.Add("field"+fmt.Sprint(i), i, false, true)
		kvs = kvs.Add("tag"+fmt.Sprint(i), fmt.Sprint(i), true, true)
	}
	return
}

func testFunc_2() {
	var kvs KVs
	for i := 0; i < lenKV; i++ {
		kvs.AddV2("field"+fmt.Sprint(i), i, false, true)
		kvs.AddV2("tag"+fmt.Sprint(i), fmt.Sprint(i), true, true)
	}
	return
}
