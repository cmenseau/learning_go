package grep_runner

import (
	"fmt"
	grep_engine "main/internal/engine"
	grep_file_scanner "main/internal/file_scanner"
	grep_parser "main/internal/parser"
	"os"
)

func Run(args []string) {
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

	scanner := grep_file_scanner.NewFileScanner(
		eng, req.Paths, req.Recursive, os.Stdout, os.Stderr)

	scanner.GoThroughFiles()
}
