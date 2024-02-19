package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/fatih/color"
)

func main() {

	// parse args

	if len(os.Args[1:]) < 2 {
		fmt.Errorf("Expecting more than 2 args, got %v", os.Args[1:])
	}

	var options []string
	var pattern_idx int

	var args = []string(os.Args[1:])

	for i, val := range args {
		if strings.HasPrefix(val, "-") {
			options = append(options, strings.Replace(val, "-", "", -1))
		} else {
			pattern_idx = i
			break
		}
	}

	pattern := args[pattern_idx]

	filenames := args[pattern_idx+1:]

	var case_insensitive bool
	var match_whole_line bool
	var invert_matching bool

	switch {
	case slices.Contains(options, "i"):
		case_insensitive = true
	case slices.Contains(options, "x"):
		match_whole_line = true
	case slices.Contains(options, "v"):
		invert_matching = true
	}

	// TODO refactor as a sort of pipeline to be able to use options together

	if invert_matching {
		fmt.Print(go_through_files(pattern, filenames, select_reverse))
	} else if match_whole_line {
		fmt.Println(go_through_files(pattern, filenames, select_whole_line))
	} else {
		if case_insensitive {
			fmt.Print(go_through_files(pattern, filenames, select_highlight_cis))
		} else {
			fmt.Print(go_through_files(pattern, filenames, select_highlight))
		}
	}

}

func go_through_files(
	keyword string,
	files []string,
	select_func func(string, string) string) string {

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
				line_output := select_func(keyword, line)

				if len(line_output) != 0 {
					if print_filename {
						line_output = color_magenta(filename) + color_cyan(":") + line_output
					}
					output += line_output
				}
			}
		}
		file.Close()
	}
	return output
}

/////// SELECT PATTERNS ON LINE FUNCTIONS //////////

// Grep default behavior : select lines containing keyword, and highlight keyword
func select_highlight(keyword string, line string) (res string) {
	if strings.Contains(line, keyword) {
		split := strings.Split(line, keyword)
		res = strings.Join(split, color_red(keyword)) + "\n"
	}
	return
}

// case-insensitive version
func select_highlight_cis(keyword string, line string) (res string) {
	var lower_keyword = strings.ToLower(keyword)
	var len_keyword = len(keyword)

	if strings.Contains(strings.ToLower(line), lower_keyword) {
		var idx int = strings.Index(strings.ToLower(line), lower_keyword)
		if idx != -1 {
			for idx != -1 {
				res += line[:idx] + color_red(line[idx:idx+len_keyword])
				line = line[idx+len_keyword:]

				idx = strings.Index(strings.ToLower(line), lower_keyword)
			}
			res += line + "\n"
		}
	}
	return
}

func select_reverse(keyword string, line string) (res string) {
	if !strings.Contains(line, keyword) {
		res = line + "\n"
	}
	return
}

func select_whole_line(keyword string, line string) (res string) {
	if line == keyword {
		res = line + "\n"
	}
	return
}

/////// COLOR OUTPUT FUNCTIONS ///////////

func color_red(str string) string {
	red := color.New(color.FgRed).SprintFunc()
	return red(str)
}

func color_magenta(str string) string {
	magenta := color.New(color.FgMagenta).SprintFunc()
	return magenta(str)
}

func color_cyan(str string) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	return cyan(str)
}
