package grep_engine

import (
	grep_line_prefix_control "main/internal/line_prefix_control"
	grep_line_select "main/internal/line_select"
	grep_output_control "main/internal/output_control"
	grep_parser "main/internal/parser"
	"testing"
)

func TestEngineLine(test *testing.T) {
	subtests := []struct {
		request  grep_parser.GrepRequest
		line     string
		filename string
		line_out string
	}{
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "abcdef\n",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "def",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{CaseInsensitive: true},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "aBcdef",
			filename: "1.txt",
			line_out: "aBcdef\n",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{InvertMatching: true},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "def",
			filename: "1.txt",
			line_out: "def\n",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{Granularity: grep_line_select.WordGranularity},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{Granularity: grep_line_select.WordGranularity},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc def",
			filename: "1.txt",
			line_out: "abc def\n",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{Granularity: grep_line_select.LineGranularity},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{Granularity: grep_line_select.LineGranularity},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "abc\n",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{OnlyMatching: true},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "abc\n",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{OnlyMatching: true},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "def",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{CountLines: true},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{FilesWithoutMatch: true},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{FilesWithMatch: true},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true},
			},
			line:     "abcdef",
			filename: "2.txt",
			line_out: "2.txt:abcdef\n",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{CaseInsensitive: true, OnlyMatching: true, Granularity: grep_line_select.WordGranularity},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true},
			},
			line:     "what aBc def",
			filename: "file.txt",
			line_out: "file.txt:aBc\n",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{OnlyMatching: true},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			line:     "abc xyz abc",
			filename: "file.txt",
			line_out: "abc\nabc\n",
		},
	}

	for _, subtest := range subtests {
		engine := Engine{&subtest.request}

		var out = engine.OutputOnLine(subtest.line, subtest.filename)

		if out != subtest.line_out {
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
		request  grep_parser.GrepRequest
		content  []lines
		filename string
		line_out string
	}{
		{
			content: []lines{},
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{CaseInsensitive: true},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{InvertMatching: true},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{Granularity: grep_line_select.AllGranularity},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{Granularity: grep_line_select.WordGranularity},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{Granularity: grep_line_select.LineGranularity},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{CountLines: true},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
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
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{CountLines: true},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true},
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
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{FilesWithoutMatch: true},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true},
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
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{FilesWithoutMatch: true},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
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
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{FilesWithMatch: true},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{},
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
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{FilesWithMatch: true},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true},
			},
			filename: "whatever.txt",
			line_out: "",
		},
		{
			request: grep_parser.GrepRequest{
				Pattern:    "abc",
				Search:     grep_line_select.SearchInfo{},
				FileOutput: grep_output_control.FileOutputRequest{},
				LinePrefix: grep_line_prefix_control.LinePrefixRequest{WithFilename: true},
			},
			filename: "whatever.txt",
			line_out: "", // not "whatever.txt:" because this is prefix on line-level
		},
	}

	for _, subtest := range subtests {
		engine := Engine{&subtest.request}

		for _, content := range subtest.content {
			engine.OutputOnLine(content.line, content.filename)
		}

		var out = engine.OutputOnWholeFile(subtest.filename)

		if out != subtest.line_out {
			test.Errorf("wanted %#v (for filename=%s req=%+v),\ngot %#v",
				subtest.line_out,
				subtest.filename, subtest.request,
				out)
		}
	}
}
