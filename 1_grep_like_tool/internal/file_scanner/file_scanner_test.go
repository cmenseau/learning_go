package grep_file_scanner

import (
	"testing"
)

func TestFileScanner(test *testing.T) {
	var aol = func(line string, filename string) string {
		return line + ":" + filename
	}

	var subtests = []struct {
		keyword      string
		files        []string
		actionOnLine func(line string, filename string) string
		line_out     string
	}{
		{
			keyword:      "good",
			files:        []string{"./test_material/test1.txt", "./test_material/test2.txt"},
			actionOnLine: aol,
			line_out: `a:./test_material/test1.txt
b:./test_material/test1.txt
c:./test_material/test1.txt
hello:./test_material/test2.txt
hi:./test_material/test2.txt
good afternoon:./test_material/test2.txt
`,
		},
		{
			keyword:      "whatever",
			files:        []string{"./test_material/test2.txt"},
			actionOnLine: aol,
			line_out: `hello:./test_material/test2.txt
hi:./test_material/test2.txt
good afternoon:./test_material/test2.txt
`,
		},
	}

	for _, subtest := range subtests {
		var out = GoThroughFiles(subtest.keyword, subtest.files, subtest.actionOnLine)

		if out != subtest.line_out {
			test.Errorf("wanted %#v (\"%v\" in : %v),\ngot %#v",
				subtest.line_out,
				subtest.keyword, subtest.files,
				out)
		}
	}
}
