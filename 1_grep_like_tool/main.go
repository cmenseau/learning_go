package main

import (
	"fmt"
	grep_engine "main/internal/engine"
	grep_file_scanner "main/internal/file_scanner"
	grep_parser "main/internal/parser"
	"os"
)

func run(args []string) {
	req, err := grep_parser.ParseArgs(args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	eng := grep_engine.Engine{Request: &req}

	scanner := grep_file_scanner.FileScanner{
		Finder:    eng,
		Paths:     req.Paths,
		Recursive: req.Recursive,
	}

	fmt.Print(scanner.GoThroughFiles())
}

func main() {
	run(os.Args[1:])
}
