package grep_file_scanner

import (
	"testing"
)

func TestFileScanner(test *testing.T) {
	var aol = func(line string, filename string) string {
		return line + ":" + filename + "\n"
	}

	var aowf = func(filename string) string {
		return "FILE:" + filename + "\n"
	}

	var subtests = []struct {
		files             []string
		actionOnLine      func(line string, filename string) string
		actionOnWholeFile func(filename string) string
		line_out          string
	}{
		{
			files:             []string{"./test_material/test1.txt", "./test_material/test2.txt"},
			actionOnLine:      aol,
			actionOnWholeFile: aowf,
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
			files:             []string{"./test_material/test2.txt"},
			actionOnLine:      aol,
			actionOnWholeFile: aowf,
			line_out: `hello:./test_material/test2.txt
hi:./test_material/test2.txt
good afternoon:./test_material/test2.txt
FILE:./test_material/test2.txt
`,
		},
	}

	for _, subtest := range subtests {
		var out = GoThroughFiles(subtest.files, subtest.actionOnLine, subtest.actionOnWholeFile)

		if out != subtest.line_out {
			test.Errorf("wanted %#v (files=%v),\ngot %#v",
				subtest.line_out,
				subtest.files,
				out)
		}
	}
}
