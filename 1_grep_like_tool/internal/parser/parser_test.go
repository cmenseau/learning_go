package grep_parser

import (
	grep_line_select "main/internal/line_select"
	grep_output_control "main/internal/output_control"
	"slices"
	"testing"
)

func TestParser(test *testing.T) {
	var subtests = []struct {
		args      []string
		pattern   string
		filenames []string
		search    grep_line_select.SearchInfo
		req       grep_output_control.OutputControlRequest
	}{
		{
			args:      []string{"lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
		},
		{
			args:      []string{"-i", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{CaseInsensitive: true},
		},
		{
			args:      []string{"-ix", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"},
		},
		{
			args:      []string{"-w", "lo", "a.txt", "b.txt", "c.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt", "b.txt", "c.txt"},
			search:    grep_line_select.SearchInfo{MatchGranularity: "word"},
		},
		{
			args:      []string{"-iwx", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"},
		},
		{
			args:      []string{"-ixw", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"},
		},
		{
			args:      []string{"-wv", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{InvertMatching: true, MatchGranularity: "word"},
		},
		{
			args:      []string{"-ixv", "whatever", "a.txt"},
			pattern:   "whatever",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line", InvertMatching: true},
		},
		{
			args:      []string{"lo", "-c", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			req:       grep_output_control.OutputControlRequest{CountLines: true},
		},
		{
			args:      []string{"lo", "-L", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			req:       grep_output_control.OutputControlRequest{FilesWithoutMatch: true},
		},
		{
			args:      []string{"lo", "-l", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			req:       grep_output_control.OutputControlRequest{FilesWithMatch: true},
		},
		{
			args:      []string{"lo", "-o", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			req:       grep_output_control.OutputControlRequest{OnlyMatching: true},
		},
	}

	for _, subtest := range subtests {
		var pattern_out, filenames_out, search_out, req_out = ParseArgs(subtest.args)

		if pattern_out != subtest.pattern ||
			!slices.Equal(filenames_out, subtest.filenames) ||
			search_out != subtest.search ||
			req_out != subtest.req {
			test.Errorf("for input \"%#v\"\nwanted: %#v, %#v, %#v, %#v\n   got: %#v, %#v, %#v, %#v",
				subtest.args,
				subtest.pattern, subtest.filenames, subtest.search, subtest.req,
				pattern_out, filenames_out, search_out, req_out)
		}
	}
}
