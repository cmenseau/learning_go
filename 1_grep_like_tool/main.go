package main

import (
	"fmt"
	grep_file_scanner "main/internal/file_scanner"
	grep_parser "main/internal/parser"
	"os"
)

func main() {

	pattern, filenames, search, _ := grep_parser.ParseArgs(os.Args[1:])

	fmt.Print(grep_file_scanner.GoThroughFiles(pattern, filenames, search))

	// general_output_control returned by ParseArgs
	// ParseArgs generates print_filename bool automatically
	// line_selector = grep_line_select.GetOutputLine(keyword, search)
	// output_controler = grep_output_control.PrintOutput(general_output_control)
	// string = grep_file_scanner.GoThroughFiles(filenames, line_selector, output_controler)
}
