package engine

import (
	"fmt"
	"main/internal/file_output"
	"main/internal/line_output"
	"main/internal/line_prefix_output"
)

type Request struct {
	Pattern    string
	Paths      []string
	Recursive  bool
	Search     line_output.SearchInfo
	FileOutput file_output.FileOutputRequest
	LinePrefix line_prefix_output.LinePrefixRequest
}

type Engine struct {
	Request      *Request
	LineSelector line_output.LineSelector
}

func NewEngine(req *Request) (Engine, error) {
	var e Engine
	e.Request = req
	ls, err := line_output.NewLineSelector(req.Pattern, req.Search)

	if err != nil {
		return Engine{}, fmt.Errorf("create engine : %w", err)
	}

	e.LineSelector = ls
	return e, nil
}

func (e Engine) OutputOnLine(line string, filename string) string {

	line_output := e.LineSelector.GetOutputLine(line)

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
