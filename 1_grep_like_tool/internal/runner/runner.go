package grep_runner

import (
	"fmt"
	grep_engine "main/internal/engine"
	grep_file_scanner "main/internal/file_scanner"
	grep_parser "main/internal/parser"
	"os"
)

func Run(args []string, out *os.File) {
	req, err := grep_parser.ParseArgs(args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	eng, err := grep_engine.NewEngine(&req)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	scanner := grep_file_scanner.FileScanner{
		Finder:    eng,
		Paths:     req.Paths,
		Recursive: req.Recursive,
	}

	// TODO : GoThroughFiles : give out to write line 1 by 1
	output, err := scanner.GoThroughFiles()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprint(out, output)
	}
}
