package grep_engine

import (
	grep_parser "main/internal/parser"

	grep_line_select "main/internal/line_select"
)

type Engine struct {
	Request *grep_parser.GrepRequest
}

func (e Engine) OutputOnLine(line string, filename string) string {

	line_output := grep_line_select.GetOutputLine(e.Request.Pattern, line, e.Request.Search)

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
