package grep_engine

import (
	grep_line_prefix_control "main/internal/line_prefix_control"
	grep_line_select "main/internal/line_select"
	grep_output_control "main/internal/output_control"
	"slices"
)

type Request struct {
	Pattern    string
	Filenames  []string
	Search     grep_line_select.SearchInfo
	FileOutput grep_output_control.FileOutputRequest
	LinePrefix grep_line_prefix_control.LinePrefixRequest
}

func (req Request) Equal(req2 Request) bool {
	return req.Pattern == req2.Pattern &&
		slices.Equal(req.Filenames, req2.Filenames) &&
		req.Search == req2.Search &&
		req.FileOutput.Equal(req2.FileOutput) &&
		req.LinePrefix == req2.LinePrefix
}

type Engine struct {
	Request *Request
}

func (e Engine) OutputOnLine(line string, filename string) string {

	line_output := grep_line_select.GetOutputLine(e.Request.Pattern, line, e.Request.Search)

	if e.Request.FileOutput.SuppressNormalOutput() {
		e.Request.FileOutput.ProcessOutputLine(line_output, filename)
		line_output = ""
	}

	if len(line_output) != 0 {
		prefix := e.Request.LinePrefix.GetPrefix(filename)

		line_output = prefix + line_output + "\n"
	}
	return line_output
}

func (e Engine) OutputOnWholeFile(filename string) string {

	line_output := e.Request.FileOutput.GetFinalOutputControl(filename)

	if len(line_output) > 0 {
		prefix := e.Request.LinePrefix.GetPrefix(filename)

		line_output = prefix + line_output + "\n"
	}

	return line_output
}
