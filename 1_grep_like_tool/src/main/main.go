package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	pattern, filenames, search := parse_args(os.Args[1:])

	fmt.Print(go_through_files(pattern, filenames, search))
}

func go_through_files(
	keyword string,
	files []string,
	search search_info) string {

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

				line_output := get_output_line(keyword, line, search)

				if len(line_output) != 0 {
					if print_filename {
						line_output = color_magenta(filename) + color_cyan(":") + line_output
					}
					output += line_output + "\n"
				}
			}
		}
		file.Close()
	}
	return output
}
