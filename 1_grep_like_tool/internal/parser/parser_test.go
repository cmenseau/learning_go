package grep_parser

import (
	grep_line_select "main/internal/line_select"
	"slices"
	"testing"
)

func TestParser(test *testing.T) {
	var subtests = []struct {
		args      []string
		pattern   string
		filenames []string
		search    grep_line_select.SearchInfo
	}{
		{
			args:      []string{"lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
		},
		{
			args:      []string{"-i", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{CaseInsensitive: true},
		},
		{
			args:      []string{"-ix", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"},
		},
		{
			args:      []string{"-w", "lo", "a.txt", "b.txt", "c.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt", "b.txt", "c.txt"},
			search:    grep_line_select.SearchInfo{MatchGranularity: "word"},
		},
		{
			args:      []string{"-iwx", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"},
		},
		{
			args:      []string{"-ixw", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line"},
		},
		{
			args:      []string{"-wv", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{InvertMatching: true, MatchGranularity: "word"},
		},
		{
			args:      []string{"-ixv", "whatever", "a.txt"},
			pattern:   "whatever",
			filenames: []string{"a.txt"},
			search:    grep_line_select.SearchInfo{CaseInsensitive: true, MatchGranularity: "line", InvertMatching: true},
		},
	}

	for _, subtest := range subtests {
		//TODO use last var
		var pattern_out, filenames_out, search_out, _ = ParseArgs(subtest.args)

		if pattern_out != subtest.pattern ||
			!slices.Equal(filenames_out, subtest.filenames) ||
			search_out != subtest.search {
			test.Errorf("for input \"%#v\"\nwanted: %#v, %#v, %#v\n   got: %#v, %#v, %#v",
				subtest.args,
				subtest.pattern, subtest.filenames, subtest.search,
				pattern_out, filenames_out, search_out)
		}
	}
}
