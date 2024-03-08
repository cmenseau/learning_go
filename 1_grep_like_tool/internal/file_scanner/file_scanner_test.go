package grep_file_scanner

import (
	grep_line_select "main/internal/line_select"
	"testing"
)

func TestFileScanner(test *testing.T) {
	var subtests = []struct {
		keyword  string
		files    []string
		search   grep_line_select.SearchInfo
		line_out string
	}{
		{
			keyword:  "good",
			files:    []string{"./test_material/test1.txt", "./test_material/test2.txt"},
			search:   grep_line_select.SearchInfo{},
			line_out: "./test_material/test2.txt:good morning\n./test_material/test2.txt:good afternoon\n./test_material/test2.txt:good evening\n",
		},
		{
			keyword:  "good",
			files:    []string{"./test_material/test2.txt"},
			search:   grep_line_select.SearchInfo{},
			line_out: "good morning\ngood afternoon\ngood evening\n",
		},
		{
			keyword:  "a",
			files:    []string{"./test_material/test1.txt", "./test_material/test2.txt"},
			search:   grep_line_select.SearchInfo{MatchGranularity: "word"},
			line_out: "./test_material/test1.txt:a\n",
		},
		{
			keyword:  "wwwww",
			files:    []string{"./test_material/test1.txt", "./test_material/test2.txt"},
			search:   grep_line_select.SearchInfo{},
			line_out: "",
		},
		{
			keyword:  "a",
			files:    []string{"./test_material/test1.txt", "./test_material/test2.txt"},
			search:   grep_line_select.SearchInfo{OnlyMatching: true},
			line_out: "./test_material/test1.txt:a\n./test_material/test2.txt:a\n",
		},
	}

	for _, subtest := range subtests {
		var out = GoThroughFiles(subtest.keyword, subtest.files, subtest.search)

		if out != subtest.line_out {
			test.Errorf("wanted %#v (\"%v\" in : %v, %v), got %#v",
				subtest.line_out, subtest.keyword,
				subtest.files, subtest.search,
				out)
		}
	}
}
