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
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt"}},
		},
		{
			args: []string{"-i", "lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}}},
		{
			args: []string{"-ix", "lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"}}},
		{
			args: []string{"-w", "lo", "a.txt", "b.txt", "c.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt", "b.txt", "c.txt"}, Search: grep_line_select.SearchInfo{MatchGranularity: "word"}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-iwx", "lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"}}},
		{
			args: []string{"-ixw", "lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"}}},
		{
			args: []string{"-wv", "lo", "a.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{InvertMatching: true, MatchGranularity: "word"}}},
		{
			args: []string{"-ixv", "whatever", "a.txt"},
			req:  GrepRequest{Pattern: "whatever", Filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line", InvertMatching: true}}},
		{
			args: []string{"lo", "-c", "a.txt"}, // also okay with pattern first
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt"}, FileOutput: grep_output_control.FileOutputRequest{CountLines: true}},
		},
		{
			args: []string{"lo", "-L", "a.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt"}, FileOutput: grep_output_control.FileOutputRequest{FilesWithoutMatch: true}},
		},
		{
			args: []string{"lo", "-l", "a.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt"}, FileOutput: grep_output_control.FileOutputRequest{FilesWithMatch: true}},
		},
		{
			args: []string{"lo", "-o", "a.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"a.txt"}, Search: grep_line_select.SearchInfo{OnlyMatching: true}},
		},
		{
			args: []string{"-H", "lo", "c.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"c.txt"}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-Hi", "lo", "c.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"c.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-H", "-i", "lo", "c.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"c.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"lo", "-Hi", "c.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"c.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"lo", "-H", "-i", "c.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"c.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-H", "lo", "-i", "c.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"c.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-lH", "lo", "c.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"c.txt"}, FileOutput: grep_output_control.FileOutputRequest{FilesWithMatch: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: false}},
		},
		{
			args: []string{"-LH", "lo", "c.txt"},
			req:  GrepRequest{Pattern: "lo", Filenames: []string{"c.txt"}, FileOutput: grep_output_control.FileOutputRequest{FilesWithoutMatch: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: false}},
		},
	}

	for _, subtest := range subtests {
		var req_out = ParseArgs(subtest.args)

		if !req_out.Equal(subtest.req) {
			test.Errorf("for input \"%#v\"\nwanted: %#v\n   got: %#v",
				subtest.args,
				subtest.req,
				req_out)
		}
	}
}
