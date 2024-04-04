package grep_parser

import (
	grep_engine "main/internal/engine"
	grep_line_prefix_control "main/internal/line_prefix_control"
	grep_line_select "main/internal/line_select"
	grep_output_control "main/internal/output_control"
	"testing"
)

func TestParser(test *testing.T) {
	var subtests = []struct {
		args []string
		req  grep_engine.Request
	}{
		{
			args: []string{"lo", "a.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt"}},
		},
		{
			args: []string{"-i", "lo", "a.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}},
		},
		{
			args: []string{"-ix", "lo", "a.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, Granularity: grep_line_select.LineGranularity}},
		},
		{
			args: []string{"-w", "lo", "a.txt", "b.txt", "c.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt", "b.txt", "c.txt"}, Search: grep_line_select.SearchInfo{Granularity: grep_line_select.WordGranularity}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-wr", "lo", "a.txt", "b.txt", "c.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt", "b.txt", "c.txt"}, Recursive: true, Search: grep_line_select.SearchInfo{Granularity: grep_line_select.WordGranularity}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-wr", "lo", "folder1", "folder2"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"folder1", "folder2"}, Recursive: true, Search: grep_line_select.SearchInfo{Granularity: grep_line_select.WordGranularity}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-iwx", "lo", "a.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, Granularity: grep_line_select.LineGranularity}},
		},
		{
			args: []string{"-ixw", "lo", "a.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, Granularity: grep_line_select.LineGranularity}},
		},
		{
			args: []string{"-wvr", "lo", "a.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Recursive: true, Search: grep_line_select.SearchInfo{InvertMatching: true, Granularity: grep_line_select.WordGranularity}},
		},
		{
			args: []string{"-wvr", "lo", "path/to/folder"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"path/to/folder"}, Recursive: true, Search: grep_line_select.SearchInfo{InvertMatching: true, Granularity: grep_line_select.WordGranularity}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-wvr", "lo", "."},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"."}, Recursive: true, Search: grep_line_select.SearchInfo{InvertMatching: true, Granularity: grep_line_select.WordGranularity}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-ixv", "whatever", "a.txt"},
			req:  grep_engine.Request{Pattern: "whatever", Paths: []string{"a.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true, Granularity: grep_line_select.LineGranularity, InvertMatching: true}},
		},
		{
			args: []string{"lo", "-c", "a.txt"}, // also okay with pattern first
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, FileOutput: grep_output_control.FileOutputRequest{CountLines: true}},
		},
		{
			args: []string{"lo", "-L", "a.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, FileOutput: grep_output_control.FileOutputRequest{FilesWithoutMatch: true}},
		},
		{
			args: []string{"lo", "-l", "a.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, FileOutput: grep_output_control.FileOutputRequest{FilesWithMatch: true}},
		},
		{
			args: []string{"lo", "-o", "a.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Search: grep_line_select.SearchInfo{OnlyMatching: true}},
		},
		{
			args: []string{"-H", "lo", "c.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-Hi", "lo", "c.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-H", "-i", "lo", "c.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"lo", "-Hi", "c.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"lo", "-H", "-i", "c.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-H", "lo", "-i", "c.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, Search: grep_line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-lH", "lo", "c.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, FileOutput: grep_output_control.FileOutputRequest{FilesWithMatch: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: false}},
		},
		{
			args: []string{"-LH", "lo", "c.txt"},
			req:  grep_engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, FileOutput: grep_output_control.FileOutputRequest{FilesWithoutMatch: true}, LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: false}},
		},
	}

	for _, subtest := range subtests {
		var req_out, err_out = ParseArgs(subtest.args)

		if !req_out.Equal(subtest.req) && err_out == nil {
			test.Errorf("for input \"%#v\"\nwanted: %#v\n   got: %#v",
				subtest.args,
				subtest.req,
				req_out)
		}
	}
}

func TestParserWrongInput(test *testing.T) {
	var subtests = []struct {
		args []string
		req  grep_engine.Request
	}{
		{
			args: []string{"a.txt"},
			req:  grep_engine.Request{},
		},
		{
			args: []string{"-i", "a.txt"},
			req:  grep_engine.Request{},
		},
	}

	for _, subtest := range subtests {
		var req_out, err_out = ParseArgs(subtest.args)

		if !req_out.Equal(subtest.req) && err_out != nil {
			test.Errorf("for input \"%#v\"\nwanted: %#v\n   got: %#v",
				subtest.args,
				subtest.req,
				req_out)
		}
	}
}
