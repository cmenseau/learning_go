package grep_file_scanner

import (
	"testing"
)

type EngineMock struct{}

func (e EngineMock) OutputOnLine(line string, filename string) string {
	return line + ":" + filename + "\n"
}

func (e EngineMock) OutputOnWholeFile(filename string) string {
	return "FILE:" + filename + "\n"
}

func TestFileScanner(test *testing.T) {

	var subtests = []struct {
		files     []string
		recursive bool
		line_out  string
	}{
		{
			files: []string{"./test_material/test1.txt", "./test_material/test2.txt"},
			line_out: `a:./test_material/test1.txt
b:./test_material/test1.txt
c:./test_material/test1.txt
FILE:./test_material/test1.txt
hello:./test_material/test2.txt
hi:./test_material/test2.txt
good afternoon:./test_material/test2.txt
FILE:./test_material/test2.txt
`,
		},
		{
			files: []string{"./test_material/test2.txt"},
			line_out: `hello:./test_material/test2.txt
hi:./test_material/test2.txt
good afternoon:./test_material/test2.txt
FILE:./test_material/test2.txt
`,
		},
		{
			files:     []string{"./test_material"},
			recursive: true,
			line_out: `a line:test_material/inner/innertest.txt
another line:test_material/inner/innertest.txt
FILE:test_material/inner/innertest.txt
a:test_material/test1.txt
b:test_material/test1.txt
c:test_material/test1.txt
FILE:test_material/test1.txt
hello:test_material/test2.txt
hi:test_material/test2.txt
good afternoon:test_material/test2.txt
FILE:test_material/test2.txt
`,
		},
		{
			files:     []string{"./test_material/inner"},
			recursive: true,
			line_out: `a line:test_material/inner/innertest.txt
another line:test_material/inner/innertest.txt
FILE:test_material/inner/innertest.txt
`,
		},
		{
			files:     []string{"./test_material/test2.txt"},
			recursive: true,
			line_out: `hello:./test_material/test2.txt
hi:./test_material/test2.txt
good afternoon:./test_material/test2.txt
FILE:./test_material/test2.txt
`,
		},
	}

	for _, subtest := range subtests {
		scanner := FileScanner{
			Finder:    EngineMock{},
			Paths:     subtest.files,
			Recursive: subtest.recursive,
		}
		var out = scanner.GoThroughFiles()

		if out != subtest.line_out {
			test.Errorf("wanted %#v (files=%v, recursive=%t),\ngot %#v",
				subtest.line_out,
				subtest.files,
				subtest.recursive,
				out)
		}
	}
}
