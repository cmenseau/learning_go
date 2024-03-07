package grep_parser

import (
	"fmt"
	grep_line_select "main/internal/line_select"
	grep_output_control "main/internal/output_control"
	grep_output_line_prefix_control "main/internal/output_line_prefix_control"
	"os"
	"slices"
	"strings"
)

type GrepRequest struct {
	Pattern       string
	Filenames     []string
	Search        grep_line_select.SearchInfo
	OutputCtl     grep_output_control.OutputControlRequest
	LinePrefixCtl grep_output_line_prefix_control.OutputLinePrefixControlRequest
}

func (req GrepRequest) Equal(req2 GrepRequest) bool {
	return req.Pattern == req2.Pattern &&
			slices.Equal(req.Filenames, req2.Filenames) &&
			req.Search == req2.Search &&
			req.OutputCtl == req2.OutputCtl &&
			req.LinePrefixCtl == req2.LinePrefixCtl
}

// parse args and returns (keyword string,	filenames []string,	search grep_line_select.SearchInfo)
func ParseArgs(args []string) (req GrepRequest) {

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

	req.Filenames = args[pattern_idx+1:]

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
	return
}
