package grep_file_scanner

import (
	"strings"
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

		var out strings.Builder
		scanner := NewFileScanner(EngineMock{}, subtest.files, subtest.recursive, &out, &out)
		scanner.GoThroughFiles()

		if out.String() != subtest.line_out {
			test.Errorf("wanted %#v (files=%v, recursive=%t),\ngot %#v",
				subtest.line_out,
				subtest.files, subtest.recursive,
				out.String())
		}
	}
}

var subtests = []struct {
	name         string
	files        []string
	recursive    bool
	combined_out string
	out          string
	err          string
}{
	{
		name:  "last-file-error",
		files: []string{"./test_material/test1.txt", "./test_material/foo.foo"},
		combined_out: `a:./test_material/test1.txt
b:./test_material/test1.txt
c:./test_material/test1.txt
FILE:./test_material/test1.txt
open ./test_material/foo.foo: no such file or directory
`,
		out: `a:./test_material/test1.txt
b:./test_material/test1.txt
c:./test_material/test1.txt
FILE:./test_material/test1.txt
`,
		err: `open ./test_material/foo.foo: no such file or directory
`,
	},
	{
		name:  "first-file-error",
		files: []string{"./test_material/foo.foo", "./test_material/test1.txt"},
		combined_out: `open ./test_material/foo.foo: no such file or directory
a:./test_material/test1.txt
b:./test_material/test1.txt
c:./test_material/test1.txt
FILE:./test_material/test1.txt
`,
		out: `a:./test_material/test1.txt
b:./test_material/test1.txt
c:./test_material/test1.txt
FILE:./test_material/test1.txt
`,
		err: `open ./test_material/foo.foo: no such file or directory
`,
	},
	{
		name:      "last-folder-error-recursive",
		files:     []string{"./test_material/test1.txt", "./test_material/foo_folder"},
		recursive: true,
		combined_out: `a:./test_material/test1.txt
b:./test_material/test1.txt
c:./test_material/test1.txt
FILE:./test_material/test1.txt
lstat ./test_material/foo_folder: no such file or directory
`,
		out: `a:./test_material/test1.txt
b:./test_material/test1.txt
c:./test_material/test1.txt
FILE:./test_material/test1.txt
`,
		err: `lstat ./test_material/foo_folder: no such file or directory
`,
	},
	{
		name:      "first-folder-error-recursive",
		files:     []string{"./test_material/foo_folder", "./test_material/test1.txt"},
		recursive: true,
		combined_out: `lstat ./test_material/foo_folder: no such file or directory
a:./test_material/test1.txt
b:./test_material/test1.txt
c:./test_material/test1.txt
FILE:./test_material/test1.txt
`,
		out: `a:./test_material/test1.txt
b:./test_material/test1.txt
c:./test_material/test1.txt
FILE:./test_material/test1.txt
`,
		err: `lstat ./test_material/foo_folder: no such file or directory
`,
	},
}

func TestErrorHandlingOrder(t *testing.T) {

	for _, subtest := range subtests {

		t.Run(subtest.name, func(t *testing.T) {

			var out strings.Builder
			scanner := NewFileScanner(EngineMock{}, subtest.files, subtest.recursive, &out, &out)
			scanner.GoThroughFiles()

			// check everything (output & errors) is logged in the right order

			if out.String() != subtest.combined_out {
				t.Errorf("wanted %#v (files=%v, recursive=%t),\ngot %#v",
					subtest.combined_out,
					subtest.files,
					subtest.recursive,
					out.String())
			}

			out.Reset()

			var err strings.Builder
			scanner = NewFileScanner(EngineMock{}, subtest.files, subtest.recursive, &out, &err)
			scanner.GoThroughFiles()

			// check output and error are written on the right place

			if out.String() != subtest.out && err.String() != subtest.err {
				t.Errorf("wanted out=%#v, err=%#v (files=%v, recursive=%t),\ngot out=%#v, err=%#v",
					subtest.out, subtest.err,
					subtest.files, subtest.recursive,
					out.String(), err.String())
			}
		})
	}
}
