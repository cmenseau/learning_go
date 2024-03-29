package grep_parser

import (
	"fmt"
	grep_line_prefix_control "main/internal/line_prefix_control"
	grep_line_select "main/internal/line_select"
	grep_output_control "main/internal/output_control"
	"os"
	"slices"
	"strings"
)

// TODO : move this into engine ?
type GrepRequest struct {
	Pattern    string
	Filenames  []string
	Search     grep_line_select.SearchInfo
	FileOutput grep_output_control.FileOutputRequest
	LinePrefix grep_line_prefix_control.LinePrefixRequest
}

func (req GrepRequest) Equal(req2 GrepRequest) bool {
	return req.Pattern == req2.Pattern &&
		slices.Equal(req.Filenames, req2.Filenames) &&
		req.Search == req2.Search &&
		req.FileOutput.Equal(req2.FileOutput) &&
		req.LinePrefix == req2.LinePrefix
}

// parse args and returns GrepRequest
func ParseArgs(args []string) GrepRequest {
	var req GrepRequest

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Expecting more than 2 args, got %v\n", os.Args[1:])
		os.Exit(2)
	}

	var options []string
	var pattern_idx int
	var first_filename_idx int
	var pattern_found bool

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
				break
			}
		}
	}

	// TODO : handle error when no pattern found

	req.Pattern = args[pattern_idx]

	req.Filenames = args[first_filename_idx:]

	if len(req.Filenames) > 1 || slices.Contains(options, "H") {
		req.LinePrefix.WithFilename = true
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
		req.Search.Granularity = grep_line_select.WordGranularity
	}
	if slices.Contains(options, "x") {
		req.Search.Granularity = grep_line_select.LineGranularity
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

	return req
}
