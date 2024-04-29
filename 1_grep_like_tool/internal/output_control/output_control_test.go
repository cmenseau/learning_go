package output_control

import (
	"testing"
)

func TestSuppressNormalOutput(test *testing.T) {
	subtests := []struct {
		ocr     FileOutputRequest
		exp_out bool
	}{
		{
			ocr:     FileOutputRequest{CountLines: true},
			exp_out: true,
		},
		{
			ocr:     FileOutputRequest{FilesWithoutMatch: true},
			exp_out: true,
		},
		{
			ocr:     FileOutputRequest{FilesWithMatch: true},
			exp_out: true,
		},
		{
			ocr:     FileOutputRequest{FilesWithMatch: true, CountLines: true},
			exp_out: true,
		},
		{
			ocr:     FileOutputRequest{},
			exp_out: false,
		},
	}

	for _, subtest := range subtests {
		var out = subtest.ocr.SuppressNormalOutput()

		if out != subtest.exp_out {
			test.Errorf("wanted %#v (%#v), got %#v",
				subtest.exp_out, subtest.ocr,
				out)
		}
	}
}

func TestGetOutputLine(test *testing.T) {
	type lines struct {
		line     string
		filename string
	}

	subtests := []struct {
		content  []lines
		filename string
		ocr      FileOutputRequest
		exp_out  string
	}{
		{
			content: []lines{
				{
					line:     "match",
					filename: "whatever.txt",
				},
				{
					line:     "",
					filename: "whatever.txt",
				},
				{
					line:     "match",
					filename: "whatever.txt",
				},
			},
			filename: "whatever.txt",
			ocr:      FileOutputRequest{CountLines: true},
			exp_out:  "2",
		},
		{
			content: []lines{
				{
					line:     "",
					filename: "whatever.txt",
				},
			},
			filename: "whatever.txt",
			ocr:      FileOutputRequest{CountLines: true},
			exp_out:  "0",
		},
		{
			content: []lines{
				{
					line:     "it matches",
					filename: "whatever.txt",
				},
				{
					line:     "",
					filename: "whatever.txt",
				},
				{
					line:     "keyword found",
					filename: "another.txt",
				},
			},
			filename: "whatever.txt",
			ocr:      FileOutputRequest{CountLines: true},
			exp_out:  "1",
		},
		{
			content: []lines{
				{
					line:     "",
					filename: "whatever.txt",
				},
			},
			filename: "whatever.txt",
			ocr:      FileOutputRequest{FilesWithoutMatch: true},
			exp_out:  "whatever.txt",
		},
		{
			content: []lines{
				{
					line:     "toto",
					filename: "whatever.txt",
				},
				{
					line:     "tata",
					filename: "whatever.txt",
				},
			},
			filename: "whatever.txt",
			ocr:      FileOutputRequest{FilesWithoutMatch: true},
			exp_out:  "",
		},
		{
			content: []lines{
				{
					line:     "toto",
					filename: "whatever.txt",
				},
				{
					line:     "",
					filename: "another.txt",
				},
			},
			filename: "whatever.txt",
			ocr:      FileOutputRequest{FilesWithoutMatch: true},
			exp_out:  "",
		},
		{
			content: []lines{
				{
					line:     "",
					filename: "whatever.txt",
				},
			},
			filename: "whatever.txt",
			ocr:      FileOutputRequest{FilesWithMatch: true},
			exp_out:  "",
		},
		{
			content: []lines{
				{
					line:     "toto",
					filename: "whatever.txt",
				},
				{
					line:     "tata",
					filename: "whatever.txt",
				},
			},
			filename: "whatever.txt",
			ocr:      FileOutputRequest{FilesWithMatch: true},
			exp_out:  "whatever.txt",
		},
		{
			content: []lines{
				{
					line:     "toto",
					filename: "whatever.txt",
				},
				{
					line:     "",
					filename: "another.txt",
				},
			},
			filename: "whatever.txt",
			ocr:      FileOutputRequest{FilesWithMatch: true},
			exp_out:  "whatever.txt",
		},
	}

	for _, subtest := range subtests {
		for _, content := range subtest.content {
			subtest.ocr.ProcessOutputLine(content.line, content.filename)
		}

		out := subtest.ocr.GetFinalOutputControl(subtest.filename)

		if out != subtest.exp_out {
			test.Errorf("wanted %#v (content %+v, for file %s, req %+v), got %#v",
				subtest.exp_out,
				subtest.content, subtest.filename, subtest.ocr,
				out)
		}
	}
}
