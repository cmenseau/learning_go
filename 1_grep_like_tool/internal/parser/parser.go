package parser

import (
	"errors"
	"fmt"
	"main/internal/engine"
	"main/internal/line_output"
	"os"
	"slices"
	"strings"
)

// parse args and returns GrepRequest
func ParseArgs(args []string) (req engine.Request, err error) {

	if len(args) < 2 {
		err = fmt.Errorf("expecting more than 2 args, got %v", os.Args[1:])
		return
	}

	var options []string
	var pattern_idx int
	var first_filename_idx int
	var pattern_found, filename_found bool

	for i, val := range args {
		if strings.HasPrefix(val, "-") {
			var chars = strings.Split(strings.Replace(val, "-", "", -1), "")
			options = append(options, chars...)
		} else {
			// not an option : either the pattern if it wasn't yet found,
			// otherwise the first filepath
			if !pattern_found {
				pattern_idx = i
				pattern_found = true
			} else {
				first_filename_idx = i
				filename_found = true
				break
			}
		}
	}

	if !pattern_found || !filename_found {
		err = errors.New("expecting arguments : [options...] keyword file... ")
		return
	}

	req.Pattern = args[pattern_idx]

	req.Paths = args[first_filename_idx:]

	if len(req.Paths) > 1 || slices.Contains(options, "H") {
		req.LinePrefix.WithFilename = true
	}

	// -r with only 1 filename : req.LinePrefix.WithFilename => false
	// -r with 1+ dir : req.LinePrefix.WithFilename => true
	if slices.Contains(options, "r") {
		req.Recursive = true

		if fi, err := os.Stat(req.Paths[0]); len(req.Paths) == 1 && err == nil && !fi.IsDir() {
			req.LinePrefix.WithFilename = false
		} else {
			req.LinePrefix.WithFilename = true
		}
	}

	if slices.Contains(options, "i") {
		req.Search.CaseInsensitive = true
	}
	if slices.Contains(options, "v") {
		req.Search.InvertMatching = true
	}
	if slices.Contains(options, "o") {
		req.Search.OnlyMatching = true
	}

	// if -x and -w specified, -x takes over
	if slices.Contains(options, "w") {
		req.Search.Granularity = line_output.WordGranularity
	}
	if slices.Contains(options, "x") {
		req.Search.Granularity = line_output.LineGranularity
	}

	// only one the the 3 option can be used,
	// with priority filesWithout over filesWith over CountLines
	if slices.Contains(options, "L") {
		req.FileOutput.FilesWithoutMatch = true
		req.LinePrefix.WithFilename = false
	} else if slices.Contains(options, "l") {
		req.FileOutput.FilesWithMatch = true
		req.LinePrefix.WithFilename = false
	} else if slices.Contains(options, "c") {
		req.FileOutput.CountLines = true
	}

	return
}
