package grep_file_scanner

import (
	"bufio"
	"fmt"
	grep_colors "main/internal/colors"           // TODO remove
	grep_line_select "main/internal/line_select" // TODO remove
	"os"
)

func GoThroughFiles(
	keyword string,
	files []string,
	search grep_line_select.SearchInfo) string {

	var output string
	var print_filename bool
	if len(files) > 1 {
		print_filename = true
	} else {
		print_filename = false
	}

	for _, filename := range files {

		file, err := os.Open(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open file %s because of %s", filename, err)
		} else {

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()

				line_output := grep_line_select.GetOutputLine(keyword, line, search)

				if len(line_output) != 0 {
					if print_filename {
						line_output = grep_colors.Color_magenta(filename) + grep_colors.Color_cyan(":") + line_output
					}
					output += line_output + "\n"
				}
			}
		}
		file.Close()
	}
	return output
}
