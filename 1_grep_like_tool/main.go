package main

import (
	"fmt"
	grep_engine "main/internal/engine"
	grep_file_scanner "main/internal/file_scanner"
	grep_parser "main/internal/parser"
	"os"
)

func main() {
	req := grep_parser.ParseArgs(os.Args[1:])

	eng := grep_engine.Engine{Request: &req}

	scanner := grep_file_scanner.FileScanner{
		Finder:    eng,
		Filenames: req.Filenames,
	}

	fmt.Print(scanner.GoThroughFiles())
}
