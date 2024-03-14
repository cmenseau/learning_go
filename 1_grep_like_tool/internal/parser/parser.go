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

type GrepRequest struct {
	Pattern    string
	filenames  []string
	Search     grep_line_select.SearchInfo
	FileOutput grep_output_control.FileOutputRequest
	LinePrefix grep_line_prefix_control.LinePrefixRequest
}

func (req GrepRequest) Equal(req2 GrepRequest) bool {
	return req.Pattern == req2.Pattern &&
		slices.Equal(req.filenames, req2.filenames) &&
		req.Search == req2.Search &&
		req.FileOutput == req2.FileOutput &&
		req.LinePrefix == req2.LinePrefix
}

// parse args and returns GrepRequest
func ParseArgs(args []string) (GrepRequest, []string) {
	var req GrepRequest

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

	req.Pattern = args[pattern_idx]

	req.filenames = args[pattern_idx+1:]

	if len(req.filenames) > 1 {
		req.LinePrefix.WithFilename = true
	}

	if slices.Contains(options, "i") {
		req.Search.CaseInsensitive = true
	}
	if slices.Contains(options, "v") {
		req.Search.InvertMatching = true
	}

	// if -x and -w specified, -x takes over
	if slices.Contains(options, "w") {
		req.Search.MatchGranularity = "word"
	}
	if slices.Contains(options, "x") {
		req.Search.MatchGranularity = "line"
	}
	return req, req.filenames
}
