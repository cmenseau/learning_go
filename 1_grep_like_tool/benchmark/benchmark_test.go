package benchmark

import (
	grep_runner "main/internal/runner"
	"os"
	"testing"
)

// go test benchmark/benchmark_test.go -bench=. -benchmem > out.txt
func BenchmarkRecursive(b *testing.B) {
	params := []string{"-iwr", `defer\|func`, "/home/menseau/Documents/Go/learning_go/0_go_class_matt_holiday"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		grep_runner.Run(params, os.Stdout)
	}
}
