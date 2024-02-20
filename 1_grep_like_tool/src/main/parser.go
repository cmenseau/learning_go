package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// parse args and returns (keyword string,	filenames []string,	search search_info)
func parse_args(args []string) (pattern string, filenames []string, search search_info) {

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
		search.case_insensitive = true
	}
	if slices.Contains(options, "v") {
		search.invert_matching = true
	}

	// if -x and -w specified, -x takes over
	if slices.Contains(options, "w") {
		search.match_granularity = "word"
	}
	if slices.Contains(options, "x") {
		search.match_granularity = "line"
	}
	return
}
