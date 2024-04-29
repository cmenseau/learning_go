package engine

import (
	"fmt"
	"main/internal/line_prefix_control"
	"main/internal/line_select"
	"main/internal/output_control"
)

type Request struct {
	Pattern    string
	Paths      []string
	Recursive  bool
	Search     line_select.SearchInfo
	FileOutput output_control.FileOutputRequest
	LinePrefix line_prefix_control.LinePrefixRequest
}

type Engine struct {
	Request      *Request
	LineSelector line_select.LineSelector
}

func NewEngine(req *Request) (Engine, error) {
	var e Engine
	e.Request = req
	ls, err := line_select.NewLineSelector(req.Pattern, req.Search)

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
