package parser

import (
	"main/internal/engine"
	"main/internal/file_output"
	"main/internal/line_output"
	"main/internal/line_prefix_output"
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
			req:  engine.Request{Paths: []string{"a.txt"}, Search: line_output.SearchInfo{Pattern: "lo"}},
		},
		{
			args: []string{"-i", "lo", "a.txt"},
			req:  engine.Request{Paths: []string{"a.txt"}, Search: line_output.SearchInfo{Pattern: "lo", CaseInsensitive: true}},
		},
		{
			args: []string{"-ix", "lo", "a.txt"},
			req:  engine.Request{Paths: []string{"a.txt"}, Search: line_output.SearchInfo{Pattern: "lo", CaseInsensitive: true, Granularity: line_output.LineGranularity}},
		},
		{
			args: []string{"-w", "lo", "a.txt", "b.txt", "c.txt"},
			req:  engine.Request{Paths: []string{"a.txt", "b.txt", "c.txt"}, Search: line_output.SearchInfo{Pattern: "lo", Granularity: line_output.WordGranularity}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-wr", "lo", "test_material/recursive/folder1/a.txt", "test_material/recursive/folder2/b.txt"},
			req:  engine.Request{Paths: []string{"test_material/recursive/folder1/a.txt", "test_material/recursive/folder2/b.txt"}, Recursive: true, Search: line_output.SearchInfo{Pattern: "lo", Granularity: line_output.WordGranularity}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-wr", "lo", "test_material/recursive/folder1", "test_material/recursive/folder2"},
			req:  engine.Request{Paths: []string{"test_material/recursive/folder1", "test_material/recursive/folder2"}, Recursive: true, Search: line_output.SearchInfo{Pattern: "lo", Granularity: line_output.WordGranularity}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-iwx", "lo", "a.txt"},
			req:  engine.Request{Paths: []string{"a.txt"}, Search: line_output.SearchInfo{Pattern: "lo", CaseInsensitive: true, Granularity: line_output.LineGranularity}},
		},
		{
			args: []string{"-ixw", "lo", "a.txt"},
			req:  engine.Request{Paths: []string{"a.txt"}, Search: line_output.SearchInfo{Pattern: "lo", CaseInsensitive: true, Granularity: line_output.LineGranularity}},
		},
		{
			args: []string{"-wvr", "lo", "test_material/recursive/folder1/a.txt"},
			req:  engine.Request{Paths: []string{"test_material/recursive/folder1/a.txt"}, Recursive: true, Search: line_output.SearchInfo{Pattern: "lo", InvertMatching: true, Granularity: line_output.WordGranularity}},
		},
		{
			args: []string{"-wvr", "lo", "test_material/recursive/folder1"},
			req:  engine.Request{Paths: []string{"test_material/recursive/folder1"}, Recursive: true, Search: line_output.SearchInfo{Pattern: "lo", InvertMatching: true, Granularity: line_output.WordGranularity}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-wvr", "lo", "."},
			req:  engine.Request{Paths: []string{"."}, Recursive: true, Search: line_output.SearchInfo{Pattern: "lo", InvertMatching: true, Granularity: line_output.WordGranularity}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-ixv", "whatever", "a.txt"},
			req:  engine.Request{Paths: []string{"a.txt"}, Search: line_output.SearchInfo{Pattern: "whatever", CaseInsensitive: true, Granularity: line_output.LineGranularity, InvertMatching: true}},
		},
		{
			args: []string{"lo", "-c", "a.txt"}, // also okay with pattern first
			req:  engine.Request{Paths: []string{"a.txt"}, Search: line_output.SearchInfo{Pattern: "lo"}, FileOutput: file_output.FileOutputRequest{CountLines: true}},
		},
		{
			args: []string{"lo", "-L", "a.txt"},
			req:  engine.Request{Paths: []string{"a.txt"}, Search: line_output.SearchInfo{Pattern: "lo"}, FileOutput: file_output.FileOutputRequest{FilesWithoutMatch: true}},
		},
		{
			args: []string{"lo", "-l", "a.txt"},
			req:  engine.Request{Paths: []string{"a.txt"}, Search: line_output.SearchInfo{Pattern: "lo"}, FileOutput: file_output.FileOutputRequest{FilesWithMatch: true}},
		},
		{
			args: []string{"lo", "-o", "a.txt"},
			req:  engine.Request{Paths: []string{"a.txt"}, Search: line_output.SearchInfo{Pattern: "lo", OnlyMatching: true}},
		},
		{
			args: []string{"-H", "lo", "c.txt"},
			req:  engine.Request{Paths: []string{"c.txt"}, Search: line_output.SearchInfo{Pattern: "lo"}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-Hi", "lo", "c.txt"},
			req:  engine.Request{Paths: []string{"c.txt"}, Search: line_output.SearchInfo{Pattern: "lo", CaseInsensitive: true}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-H", "-i", "lo", "c.txt"},
			req:  engine.Request{Paths: []string{"c.txt"}, Search: line_output.SearchInfo{Pattern: "lo", CaseInsensitive: true}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"lo", "-Hi", "c.txt"},
			req:  engine.Request{Paths: []string{"c.txt"}, Search: line_output.SearchInfo{Pattern: "lo", CaseInsensitive: true}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"lo", "-H", "-i", "c.txt"},
			req:  engine.Request{Paths: []string{"c.txt"}, Search: line_output.SearchInfo{Pattern: "lo", CaseInsensitive: true}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-H", "lo", "-i", "c.txt"},
			req:  engine.Request{Paths: []string{"c.txt"}, Search: line_output.SearchInfo{Pattern: "lo", CaseInsensitive: true}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true}},
		},
		{
			args: []string{"-lH", "lo", "c.txt"},
			req:  engine.Request{Paths: []string{"c.txt"}, Search: line_output.SearchInfo{Pattern: "lo"}, FileOutput: file_output.FileOutputRequest{FilesWithMatch: true}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: false}},
		},
		{
			args: []string{"-LH", "lo", "c.txt"},
			req:  engine.Request{Paths: []string{"c.txt"}, Search: line_output.SearchInfo{Pattern: "lo"}, FileOutput: file_output.FileOutputRequest{FilesWithoutMatch: true}, LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: false}},
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
