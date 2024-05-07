package engine

type LineSelector interface {
	GetResult(line string) string
}

type FileLevelMotor interface {
	ProcessLine(selected_line string, filename string) string
	GetFileLevelResult(filename string) string
}

type FilePrefixMotor interface {
	GetPrefix(filename string) string
}

type Engine struct {
	request *Request
	ls      LineSelector
	flm     FileLevelMotor
	fpm     FilePrefixMotor
}

func NewEngine(req *Request, ls LineSelector, flm FileLevelMotor, fpm FilePrefixMotor) (Engine, error) {
	var e Engine
	e.request = req
	e.ls = ls
	e.flm = flm
	e.fpm = fpm
	return e, nil
}

func (e Engine) OutputOnLine(line string, filename string) string {

	line_output := e.ls.GetResult(line)

	line_output = e.flm.ProcessLine(line_output, filename)

	if len(line_output) != 0 {
		prefix := e.fpm.GetPrefix(filename)

		line_output = prefix + line_output + "\n"
	}
	return line_output
}

func (e Engine) OutputOnWholeFile(filename string) string {

	line_output := e.flm.GetFileLevelResult(filename)

	if len(line_output) > 0 {
		prefix := e.fpm.GetPrefix(filename)

		line_output = prefix + line_output + "\n"
	}

	return line_output
}
