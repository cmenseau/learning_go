package benchmark

import (
	"main/internal/runner"
	"testing"
)

// go test benchmark/benchmark_test.go -bench=. -benchmem > benchmark/out/out.txt
func BenchmarkRecursive(b *testing.B) {
	params := []string{"-iwr", `defer\|func`, "/home/menseau/Documents/Go/learning_go/0_go_class_matt_holiday"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runner.Run(params)
	}
}
