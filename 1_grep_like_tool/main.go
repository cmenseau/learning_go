package main

import (
	grep_runner "main/internal/runner"
	"os"
)

func main() {
	grep_runner.Run(os.Args[1:])
}
