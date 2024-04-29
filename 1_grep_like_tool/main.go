package main

import (
	"main/internal/runner"
	"os"
)

func main() {
	runner.Run(os.Args[1:])
}
