package main

import (
	"slices"
	"testing"
)

func TestParser(test *testing.T) {
	var subtests = []struct {
		args      []string
		pattern   string
		filenames []string
		search    search_info
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
			search:    search_info{case_insensitive: true},
		},
		{
			args:      []string{"-ix", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    search_info{case_insensitive: true, match_granularity: "line"},
		},
		{
			args:      []string{"-w", "lo", "a.txt", "b.txt", "c.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt", "b.txt", "c.txt"},
			search:    search_info{match_granularity: "word"},
		},
		{
			args:      []string{"-iwx", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    search_info{case_insensitive: true, match_granularity: "line"},
		},
		{
			args:      []string{"-ixw", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    search_info{case_insensitive: true, match_granularity: "line"},
		},
		{
			args:      []string{"-wv", "lo", "a.txt"},
			pattern:   "lo",
			filenames: []string{"a.txt"},
			search:    search_info{invert_matching: true, match_granularity: "word"},
		},
		{
			args:      []string{"-ixv", "whatever", "a.txt"},
			pattern:   "whatever",
			filenames: []string{"a.txt"},
			search:    search_info{case_insensitive: true, match_granularity: "line", invert_matching: true},
		},
	}

	for _, subtest := range subtests {
		var pattern_out, filenames_out, search_out = parse_args(subtest.args)

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
