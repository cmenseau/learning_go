package main

import (
	"fmt"
	grep_colors "main/internal/colors"
	grep_file_scanner "main/internal/file_scanner"
	grep_line_select "main/internal/line_select"
	grep_parser "main/internal/parser"
	"os"
)

func main() {

	req := grep_parser.ParseArgs(os.Args[1:])

	var actionOnLine = func(line string, filename string) string {

		var print_filename bool

		// TODO : use flag from req insteqd of filenames
		if len(req.Filenames) > 1 {
			print_filename = true
		} else {
			print_filename = false
		}

		line_output := grep_line_select.GetOutputLine(req.Pattern, line, req.Search)

		if len(line_output) != 0 {
			if print_filename {
				line_output = grep_colors.Color_magenta(filename) + grep_colors.Color_cyan(":") + line_output
			}
		}
		return line_output
	}

	fmt.Print(grep_file_scanner.GoThroughFiles(req.Pattern, req.Filenames, actionOnLine))

	// greprequest returned by ParseArgs
	// ParseArgs generates print_filename bool automatically
	// line_selector = grep_line_select.GetOutputLine(keyword, search)
	// output_controler = grep_output_control.PrintOutput(general_output_control)
	// line_prefix_controler = grep_line_prefix_control.PrintOutput(line_prefix_control)
	// string = grep_file_scanner.GoThroughFiles(filenames, line_selector, output_controler)
}
