package grep_parser

import (
	"fmt"
	grep_line_select "main/internal/line_select"
	"os"
	"slices"
	"strings"
)

// parse args and returns (keyword string,	filenames []string,	search grep_line_select.SearchInfo)
func ParseArgs(args []string) (pattern string, filenames []string, search grep_line_select.SearchInfo) {

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Expecting more than 2 args, got %v\n", os.Args[1:])
		os.Exit(2)
	}

	var options []string
	var pattern_idx int

	for i, val := range args {
		if strings.HasPrefix(val, "-") {
			var chars = strings.Split(strings.Replace(val, "-", "", -1), "")
			options = append(options, chars...)
		} else {
			pattern_idx = i
			break
		}
	}

	pattern = args[pattern_idx]

	filenames = args[pattern_idx+1:]

	if slices.Contains(options, "i") {
		search.CaseInsensitive = true
	}
	if slices.Contains(options, "v") {
		search.InvertMatching = true
	}

	// if -x and -w specified, -x takes over
	if slices.Contains(options, "w") {
		search.MatchGranularity = "word"
	}
	if slices.Contains(options, "x") {
		search.MatchGranularity = "line"
	}
	return
}
