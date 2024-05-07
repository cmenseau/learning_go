package engine

import (
	"main/internal/file_output"
	"main/internal/line_output"
	"main/internal/line_prefix_output"
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
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "abcdef\n",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "def",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", CaseInsensitive: true},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "aBcdef",
			filename: "1.txt",
			line_out: "aBcdef\n",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", InvertMatching: true},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "def",
			filename: "1.txt",
			line_out: "def\n",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", Granularity: line_output.WordGranularity},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", Granularity: line_output.WordGranularity},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "abc def",
			filename: "1.txt",
			line_out: "abc def\n",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", Granularity: line_output.LineGranularity},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", Granularity: line_output.LineGranularity},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "abc\n",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", OnlyMatching: true},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "abcdef",
			filename: "1.txt",
			line_out: "abc\n",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", OnlyMatching: true},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "def",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{CountLines: true},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{FilesWithoutMatch: true},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{FilesWithMatch: true},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "abc",
			filename: "1.txt",
			line_out: "",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true},
			},
			line:     "abcdef",
			filename: "2.txt",
			line_out: "2.txt:abcdef\n",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", CaseInsensitive: true, OnlyMatching: true, Granularity: line_output.WordGranularity},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true},
			},
			line:     "what aBc def",
			filename: "file.txt",
			line_out: "file.txt:aBc\n",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", OnlyMatching: true},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			line:     "abc xyz abc",
			filename: "file.txt",
			line_out: "abc\nabc\n",
		},
	}

	for _, subtest := range subtests {

		lineSelector, err := line_output.NewLineSelector(subtest.request.Search)

		if err != nil {
			test.Fatal(err.Error())
		}

		linePrefix := line_prefix_output.LinePrefixSelector{Lpr: &subtest.request.LinePrefix}
		fileSelector := file_output.FileOutputSelector{Fo: &subtest.request.FileOutput}

		// TODO : use mock of lineSelector, fileSelector, linePrefix
		engine, err := NewEngine(lineSelector, &fileSelector, linePrefix)

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
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", CaseInsensitive: true},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", InvertMatching: true},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", Granularity: line_output.AllGranularity},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", Granularity: line_output.WordGranularity},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc", Granularity: line_output.LineGranularity},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
			},
			filename: "1.txt",
			line_out: "",
		},
		{
			content: []lines{},
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{CountLines: true},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
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
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{CountLines: true},
				LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true},
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
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{FilesWithoutMatch: true},
				LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true},
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
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{FilesWithoutMatch: true},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
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
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{FilesWithMatch: true},
				LinePrefix: line_prefix_output.LinePrefixRequest{},
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
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{FilesWithMatch: true},
				LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true},
			},
			filename: "whatever.txt",
			line_out: "",
		},
		{
			request: Request{
				Search:     line_output.SearchInfo{Pattern: "abc"},
				FileOutput: file_output.FileOutputRequest{},
				LinePrefix: line_prefix_output.LinePrefixRequest{WithFilename: true},
			},
			filename: "whatever.txt",
			line_out: "", // not "whatever.txt:" because this is prefix on line-level
		},
	}

	for _, subtest := range subtests {

		lineSelector, err := line_output.NewLineSelector(subtest.request.Search)

		if err != nil {
			test.Fatal(err.Error())
		}

		linePrefix := line_prefix_output.LinePrefixSelector{Lpr: &subtest.request.LinePrefix}

		fileSelector := file_output.FileOutputSelector{Fo: &subtest.request.FileOutput}

		// TODO : use mock of lineSelector, fileSelector, linePrefix
		engine, err := NewEngine(lineSelector, &fileSelector, linePrefix)

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
