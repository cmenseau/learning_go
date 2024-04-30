package runner

import (
	"fmt"
	"main/internal/engine"
	"main/internal/file_scanner"
	"main/internal/parser"
	"os"
)

func Run(args []string) {
	req, err := parser.ParseArgs(args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	eng, err := engine.NewEngine(&req)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	scanner := file_scanner.NewFileScanner(
		eng, req, os.Stdout, os.Stderr)

	scanner.GoThroughFiles()
}
