package runner

import (
	"fmt"
	"main/internal/engine"
	"main/internal/file_output"
	"main/internal/file_scanner"
	"main/internal/line_output"
	"main/internal/line_prefix_output"
	"main/internal/parser"
	"os"
)

func Run(args []string) {
	req, err := parser.ParseArgs(args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	lineSelector, err := line_output.NewLineSelector(req.Search)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	linePrefixSelector := line_prefix_output.LinePrefixSelector{Lpr: &req.LinePrefix}
	fileSelector := file_output.FileOutputSelector{Fo: &req.FileOutput}

	eng, err := engine.NewEngine(lineSelector, &fileSelector, linePrefixSelector)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	scanner := file_scanner.NewFileScanner(eng, req, os.Stdout, os.Stderr)

	scanner.GoThroughFiles()
}
