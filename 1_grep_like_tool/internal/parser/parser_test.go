package parser

import (
	"main/internal/engine"
	"main/internal/line_prefix_control"
	"main/internal/line_select"
	"main/internal/output_control"
	"reflect"
	"testing"
)

// TODO : add path validator in parser
// TODO : use real filepath to test

func TestParser(test *testing.T) {
	var subtests = []struct {
		args []string
		req  engine.Request
	}{
		{
			args: []string{"lo", "a.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"a.txt"}},
		},
		{
			args: []string{"-i", "lo", "a.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Search: line_select.SearchInfo{CaseInsensitive: true}},
		},
		{
			args: []string{"-ix", "lo", "a.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Search: line_select.SearchInfo{CaseInsensitive: true, Granularity: line_select.LineGranularity}},
		},
		{
			args: []string{"-w", "lo", "a.txt", "b.txt", "c.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"a.txt", "b.txt", "c.txt"}, Search: line_select.SearchInfo{Granularity: line_select.WordGranularity}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-wr", "lo", "test_material/recursive/folder1/a.txt", "test_material/recursive/folder2/b.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"test_material/recursive/folder1/a.txt", "test_material/recursive/folder2/b.txt"}, Recursive: true, Search: line_select.SearchInfo{Granularity: line_select.WordGranularity}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-wr", "lo", "test_material/recursive/folder1", "test_material/recursive/folder2"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"test_material/recursive/folder1", "test_material/recursive/folder2"}, Recursive: true, Search: line_select.SearchInfo{Granularity: line_select.WordGranularity}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-iwx", "lo", "a.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Search: line_select.SearchInfo{CaseInsensitive: true, Granularity: line_select.LineGranularity}},
		},
		{
			args: []string{"-ixw", "lo", "a.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Search: line_select.SearchInfo{CaseInsensitive: true, Granularity: line_select.LineGranularity}},
		},
		{
			args: []string{"-wvr", "lo", "test_material/recursive/folder1/a.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"test_material/recursive/folder1/a.txt"}, Recursive: true, Search: line_select.SearchInfo{InvertMatching: true, Granularity: line_select.WordGranularity}},
		},
		{
			args: []string{"-wvr", "lo", "test_material/recursive/folder1"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"test_material/recursive/folder1"}, Recursive: true, Search: line_select.SearchInfo{InvertMatching: true, Granularity: line_select.WordGranularity}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-wvr", "lo", "."},
			req:  engine.Request{Pattern: "lo", Paths: []string{"."}, Recursive: true, Search: line_select.SearchInfo{InvertMatching: true, Granularity: line_select.WordGranularity}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-ixv", "whatever", "a.txt"},
			req:  engine.Request{Pattern: "whatever", Paths: []string{"a.txt"}, Search: line_select.SearchInfo{CaseInsensitive: true, Granularity: line_select.LineGranularity, InvertMatching: true}},
		},
		{
			args: []string{"lo", "-c", "a.txt"}, // also okay with pattern first
			req:  engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, FileOutput: output_control.FileOutputRequest{CountLines: true}},
		},
		{
			args: []string{"lo", "-L", "a.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, FileOutput: output_control.FileOutputRequest{FilesWithoutMatch: true}},
		},
		{
			args: []string{"lo", "-l", "a.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, FileOutput: output_control.FileOutputRequest{FilesWithMatch: true}},
		},
		{
			args: []string{"lo", "-o", "a.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"a.txt"}, Search: line_select.SearchInfo{OnlyMatching: true}},
		},
		{
			args: []string{"-H", "lo", "c.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-Hi", "lo", "c.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, Search: line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-H", "-i", "lo", "c.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, Search: line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"lo", "-Hi", "c.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, Search: line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"lo", "-H", "-i", "c.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, Search: line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-H", "lo", "-i", "c.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, Search: line_select.SearchInfo{CaseInsensitive: true}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-lH", "lo", "c.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, FileOutput: output_control.FileOutputRequest{FilesWithMatch: true}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: false}},
		},
		{
			args: []string{"-LH", "lo", "c.txt"},
			req:  engine.Request{Pattern: "lo", Paths: []string{"c.txt"}, FileOutput: output_control.FileOutputRequest{FilesWithoutMatch: true}, LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: false}},
		},
	}

	for _, subtest := range subtests {
		var req_out, err_out = ParseArgs(subtest.args)

		if !reflect.DeepEqual(req_out, subtest.req) && err_out == nil {
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
		req  engine.Request
	}{
		{
			args: []string{"a.txt"},
			req:  engine.Request{},
		},
		{
			args: []string{"-i", "a.txt"},
			req:  engine.Request{},
		},
	}

	for _, subtest := range subtests {
		var req_out, err_out = ParseArgs(subtest.args)

		if !reflect.DeepEqual(req_out, subtest.req) && err_out != nil {
			test.Errorf("for input \"%#v\"\nwanted: %#v\n   got: %#v",
				subtest.args,
				subtest.req,
				req_out)
		}
	}
}
