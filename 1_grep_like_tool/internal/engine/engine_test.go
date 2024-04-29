package engine

import (
	"main/internal/line_prefix_control"
	"main/internal/line_select"
	"main/internal/output_control"
	"testing"
)

func TestEngineLine(test *testing.T) {
	subtests := []struct {
		request  Request
		line     string
		filename string
		line_out string
	}{
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "abcdef\n",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "def",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{CaseInsensitive: true},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "aBcdef",
			filename: "1.txt",
			line_out: "aBcdef\n",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{InvertMatching: true},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "def",
			filename: "1.txt",
			line_out: "def\n",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{Granularity: line_select.WordGranularity},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{Granularity: line_select.WordGranularity},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc def",
			filename: "1.txt",
			line_out: "abc def\n",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{Granularity: line_select.LineGranularity},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{Granularity: line_select.LineGranularity},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "abc\n",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{OnlyMatching: true},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "abc\n",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{OnlyMatching: true},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "def",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{CountLines: true},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{FilesWithoutMatch: true},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{FilesWithMatch: true},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true},
			},
			line:     "abcdef",
			filename: "2.txt",
			line_out: "2.txt:abcdef\n",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{CaseInsensitive: true, OnlyMatching: true, Granularity: line_select.WordGranularity},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true},
			},
			line:     "what aBc def",
			filename: "file.txt",
			line_out: "file.txt:aBc\n",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{OnlyMatching: true},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc xyz abc",
			filename: "file.txt",
			line_out: "abc\nabc\n",
		},
	}

	for _, subtest := range subtests {
		engine, err := NewEngine(&subtest.request)

		var out = engine.OutputOnLine(subtest.line, subtest.filename)

		if out != subtest.line_out && err == nil {
			test.Errorf("wanted %#v (for line=%s filename=%s req=%+v),\ngot %#v",
				subtest.line_out,
				subtest.line, subtest.filename, subtest.request,
				out)
		}
	}
}

func TestEngineWholeFile(test *testing.T) {
	type lines struct {
		line     string
		filename string
	}

	subtests := []struct {
		request  Request
		content  []lines
		filename string
		line_out string
	}{
		{
			content: []lines{},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{CaseInsensitive: true},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{InvertMatching: true},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{Granularity: line_select.AllGranularity},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{Granularity: line_select.WordGranularity},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{Granularity: line_select.LineGranularity},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{CountLines: true},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "0\n",
		},
		{
			content: []lines{
				{
					line:     "abc",
					filename: "1.txt",
				},
			},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{CountLines: true},
				LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true},
			},
			filename: "1.txt",
			line_out: "1.txt:1\n",
		},
		{
			content: []lines{
				{
					line:     "abc",
					filename: "1.txt",
				},
			},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{FilesWithoutMatch: true},
				LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{
				{
					line:     "def",
					filename: "1.txt",
				},
			},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{FilesWithoutMatch: true},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "1.txt\n",
		},
		{
			content: []lines{
				{
					line:     "abc",
					filename: "whatever.txt",
				},
			},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{FilesWithMatch: true},
				LinePrefix: line_prefix_control.LinePrefixRequest{},
			},
			filename: "whatever.txt",
			line_out: "whatever.txt\n",
		},
		{
			content: []lines{
				{
					line:     "def",
					filename: "whatever.txt",
				},
			},
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{FilesWithMatch: true},
				LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true},
			},
			filename: "whatever.txt",
			line_out: "",
		},
		{
			request: Request{
				Pattern:    "abc",
				Search:     line_select.SearchInfo{},
				FileOutput: output_control.FileOutputRequest{},
				LinePrefix: line_prefix_control.LinePrefixRequest{WithFilename: true},
			},
			filename: "whatever.txt",
			line_out: "", // not "whatever.txt:" because this is prefix on line-level
		},
	}

	for _, subtest := range subtests {
		engine, err := NewEngine(&subtest.request)

		for _, content := range subtest.content {
			engine.OutputOnLine(content.line, content.filename)
		}

		var out = engine.OutputOnWholeFile(subtest.filename)

		if out != subtest.line_out && err == nil {
			test.Errorf("wanted %#v (for filename=%s req=%+v),\ngot %#v",
				subtest.line_out,
				subtest.filename, subtest.request,
				out)
		}
	}
}
