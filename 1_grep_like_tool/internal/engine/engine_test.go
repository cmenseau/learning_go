package engine

import (
	"fmt"
	"testing"
)

type LineSelectorMock struct {
	returnEmpty bool
}

func (m LineSelectorMock) GetResult(line string) string {
	if m.returnEmpty {
		return ""
	}
	return "LINE=" + line
}

type FileLevelMotorMock struct {
	returnEmpty bool
}

func (m FileLevelMotorMock) ProcessLine(selected_line string, filename string) string {
	if m.returnEmpty {
		return ""
	}
	return fmt.Sprintf("PROCESSED=(%s, %s)", selected_line, filename)
}

func (m FileLevelMotorMock) GetFileLevelResult(filename string) string {
	if m.returnEmpty {
		return ""
	}
	return "RESULT=" + filename
}

type FilePrefixMotorMock struct{}

func (m FilePrefixMotorMock) GetPrefix(filename string) string {
	return fmt.Sprintf("PREFIX=%s;", filename)
}

func TestEngineLine(test *testing.T) {
	subtests := []struct {
		line     string
		filename string
		line_out string
	}{
		{
			line:     "abcdef",
			filename: "1.txt",
			line_out: "PREFIX=1.txt;PROCESSED=(LINE=abcdef, 1.txt)\n",
		},
	}

	for _, subtest := range subtests {

		mockLine := LineSelectorMock{}
		mockFile := FileLevelMotorMock{}
		mockPrefix := FilePrefixMotorMock{}

		engine, err := NewEngine(mockLine, &mockFile, mockPrefix)

		var out = engine.OutputOnLine(subtest.line, subtest.filename)

		if out != subtest.line_out && err == nil {
			test.Errorf("wanted %#v (for line=%s filename=%s),\ngot %#v",
				subtest.line_out,
				subtest.line, subtest.filename,
				out)
		}
	}
}

func TestEngineLineEmptyLineNoPrefix(test *testing.T) {
	subtests := []struct {
		line     string
		filename string
		line_out string
	}{
		{
			line:     "abcdef",
			filename: "1.txt",
			line_out: "",
		},
	}

	for _, subtest := range subtests {

		mockLine := LineSelectorMock{returnEmpty: true}
		mockFile := FileLevelMotorMock{returnEmpty: true}
		mockPrefix := FilePrefixMotorMock{}

		engine, err := NewEngine(mockLine, &mockFile, mockPrefix)

		var out = engine.OutputOnLine(subtest.line, subtest.filename)

		if out != subtest.line_out && err == nil {
			test.Errorf("wanted %#v (for line=%s filename=%s),\ngot %#v",
				subtest.line_out,
				subtest.line, subtest.filename,
				out)
		}
	}
}

func TestEngineWholeFile(test *testing.T) {
	subtests := []struct {
		filename string
		line_out string
	}{
		{
			filename: "1.txt",
			line_out: "PREFIX=1.txt;RESULT=1.txt\n",
		},
	}

	for _, subtest := range subtests {

		mockLine := LineSelectorMock{}
		mockFile := FileLevelMotorMock{}
		mockPrefix := FilePrefixMotorMock{}

		engine, err := NewEngine(mockLine, &mockFile, mockPrefix)

		var out = engine.OutputOnWholeFile(subtest.filename)

		if out != subtest.line_out && err == nil {
			test.Errorf("wanted %#v (for filename=%s),\ngot %#v",
				subtest.line_out,
				subtest.filename,
				out)
		}
	}
}

func TestEngineWholeFileEmptyResultNoPrefix(test *testing.T) {
	subtests := []struct {
		filename string
		line_out string
	}{
		{
			filename: "1.txt",
			line_out: "",
		},
	}

	for _, subtest := range subtests {

		mockLine := LineSelectorMock{}
		mockFile := FileLevelMotorMock{returnEmpty: true}
		mockPrefix := FilePrefixMotorMock{}

		engine, err := NewEngine(mockLine, &mockFile, mockPrefix)

		var out = engine.OutputOnWholeFile(subtest.filename)

		if out != subtest.line_out && err == nil {
			test.Errorf("wanted %#v (for filename=%s),\ngot %#v",
				subtest.line_out,
				subtest.filename,
				out)
		}
	}
}
