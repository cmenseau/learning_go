package main

import "testing"

// go test -bench=.
// goos: linux
// goarch: amd64
// cpu: Intel(R) Core(TM) i5-5200U CPU @ 2.20GHz
// BenchmarkFiboSeq-4      60851408                16.64 ns/op
// BenchmarkFiboRec-4         24254             48720 ns/op
// PASS
// ok      _/home/menseau/Documents/Go/learning_go/0_go_class_matt_holiday/34-35_benchmarking_mech_sympathy        2.721s

func BenchmarkFiboSeq(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacciSeq(20)
	}
}

func BenchmarkFiboRec(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacciRec(20)
	}
}
