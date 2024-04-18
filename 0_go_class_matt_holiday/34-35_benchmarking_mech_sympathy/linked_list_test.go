package main

import "testing"

// go test -bench="Linked|Slice"
// goos: linux
// goarch: amd64
// cpu: Intel(R) Core(TM) i5-5200U CPU @ 2.20GHz
// BenchmarkLinkedList-4             245796              4981 ns/op
// BenchmarkSlice-4                 2585652               469.0 ns/op
// PASS
// ok      _/home/menseau/Documents/Go/learning_go/0_go_class_matt_holiday/34-35_benchmarking_mech_sympathy        3.964s

func BenchmarkLinkedList(b *testing.B) {

	for i := 0; i < b.N; i++ {
		lst := mkList(100)
		sumList(lst)
	}

}

func BenchmarkSlice(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sl := mkSlice(100)
		sumSlice(sl)
	}

}
