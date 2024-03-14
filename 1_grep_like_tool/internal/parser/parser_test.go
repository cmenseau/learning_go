package grep_parser

import (
	grep_line_prefix_control "main/internal/line_prefix_control"
	grep_line_select "main/internal/line_select"
	grep_output_control "main/internal/output_control"
	"testing"
)

func TestParser(test *testing.T) {
	var subtests = []struct {
		args []string
		req  GrepRequest
	}{
		{
			args: []string{"lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt"}},
		},
		{
			args: []string{"-i", "lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}}},
		{
			args: []string{"-ix", "lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"}}},
		{
			args: []string{"-w", "lo", "a.txt", "b.txt", "c.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt", "b.txt", "c.txt"}, Search: grep_line_select.SearchInfo{MatchGranularity: "word"}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-iwx", "lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"}}},
		{
			args: []string{"-ixw", "lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"}}},
		{
			args: []string{"-wv", "lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{InvertMatching: true, MatchGranularity: "word"}}},
		{
			args: []string{"-ixv", "whatever", "a.txt"},
			req:  GrepRequest{Pattern: "whatever", filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line", InvertMatching: true}}},
		{
			args: []string{"lo", "-c", "a.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt"}, FileOutput: grep_output_control.FileOutputRequest{CountLines: true}},
		},
		{
			args: []string{"lo", "-L", "a.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt"}, FileOutput: grep_output_control.FileOutputRequest{FilesWithoutMatch: true}},
		},
		{
			args: []string{"lo", "-l", "a.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt"}, FileOutput: grep_output_control.FileOutputRequest{FilesWithMatch: true}},
		},
		{
			args: []string{"lo", "-o", "a.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{OnlyMatching: true}},
		},
		{
			args: []string{"-H", "lo", "c.txt"},
			req:  GrepRequest{Pattern: "lo", filenames: []string{"c.txt"}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
	}

	for _, subtest := range subtests {
		var req_out, _ = ParseArgs(subtest.args)
		//TODO update this test to check value of returned filenames

		if !req_out.Equal(subtest.req) {
			test.Errorf("for input \"%#v\"\nwanted: %#v\n   got: %#v",
				subtest.args,
				subtest.req,
				req_out)
		}
	}
}
