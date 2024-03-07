package main

import (
	"fmt"
	grep_file_scanner "main/internal/file_scanner"
	grep_parser "main/internal/parser"
	"os"
)

func main() {

	req := grep_parser.ParseArgs(os.Args[1:])

	fmt.Print(grep_file_scanner.GoThroughFiles(req.Pattern, req.Filenames, req.Search))

	// general_output_control returned by ParseArgs
	// ParseArgs generates print_filename bool automatically
	// line_selector = grep_line_select.GetOutputLine(keyword, search)
	// output_controler = grep_output_control.PrintOutput(general_output_control)
	// line_prefix_controler = grep_line_prefix_control.PrintOutput(line_prefix_control)
	// string = grep_file_scanner.GoThroughFiles(filenames, line_selector, output_controler)
}
