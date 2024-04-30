package line_output

import (
	"fmt"
	"slices"
	"testing"
)

type Subtest struct {
	search  SearchInfo
	text    string
	exp_out string
	col_out [][2]int
}

func (subtest Subtest) runST() func(t *testing.T) {

	return func(t *testing.T) {
		var ls, err_out = NewLineSelector(subtest.search)

		if err_out != nil {
			t.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.search, err_out)
		}

		out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) {
			t.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v, %v",
				subtest.exp_out, subtest.col_out,
				subtest.search, subtest.text,
				out.line, out.keywordRanges, err_out)
		}
	}
}

var testHighlight = []Subtest{
	{
		search:  SearchInfo{Pattern: "lo"},
		text:    "loutre",
		exp_out: "loutre",
		col_out: [][2]int{{0, 2}},
	},
	{
		search:  SearchInfo{Pattern: "lo"},
		text:    "Loutre",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "é"},
		text:    "canapé",
		exp_out: "canapé",
		col_out: [][2]int{{5, 7}}, // é is 2 bytes
	},
	{
		search:  SearchInfo{Pattern: "abc"},
		text:    "loutre",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "lo"},
		text:    "loutre à clou",
		exp_out: "loutre à clou",
		col_out: [][2]int{{0, 2}, {11, 13}}, // à is 2 bytes
	},
}

func TestSelectHightlight(t *testing.T) {
	for _, st := range testHighlight {
		t.Run(
			fmt.Sprintf("find %#v in %#v", st.search, st.text),
			st.runST())
	}
}

var testsHighlightInsensitive = []Subtest{
	{
		search:  SearchInfo{Pattern: "lo", CaseInsensitive: true},
		text:    "loutre",
		exp_out: "loutre",
		col_out: [][2]int{{0, 2}},
	},
	{
		search:  SearchInfo{Pattern: "lo", CaseInsensitive: true},
		text:    "Loutre",
		exp_out: "Loutre",
		col_out: [][2]int{{0, 2}},
	},
	{
		search:  SearchInfo{Pattern: "é", CaseInsensitive: true},
		text:    "canapé",
		exp_out: "canapé",
		col_out: [][2]int{{5, 7}}, // é is 2 bytes
	},
	{
		search:  SearchInfo{Pattern: "É", CaseInsensitive: true},
		text:    "canapé",
		exp_out: "canapé",
		col_out: [][2]int{{5, 7}}, // é is 2 bytes
	},
	{
		search:  SearchInfo{Pattern: "abc", CaseInsensitive: true},
		text:    "loutre",
		exp_out: "",
	},
}

func TestSelectHightlightInsensitive(t *testing.T) {
	for _, st := range testsHighlightInsensitive {
		t.Run(
			fmt.Sprintf("find %#v in %#v", st.search, st.text),
			st.runST())
	}
}

var testsReverse = []Subtest{
	{
		search:  SearchInfo{Pattern: "lo", InvertMatching: true},
		text:    "loutre",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "lo", InvertMatching: true},
		text:    "Loutre",
		exp_out: "Loutre",
	},
	{
		search:  SearchInfo{Pattern: "é", InvertMatching: true},
		text:    "canapé",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "abc", InvertMatching: true},
		text:    "loutre",
		exp_out: "loutre",
	},
}

func TestSelectReverse(t *testing.T) {

	for _, st := range testsReverse {
		t.Run(
			fmt.Sprintf("find %#v in %#v", st.search, st.text),
			st.runST())
	}
}

var testsWholeLine = []Subtest{
	{
		search:  SearchInfo{Pattern: "lo", Granularity: LineGranularity},
		text:    "loutre",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "Loutre", Granularity: LineGranularity},
		text:    "Loutre",
		exp_out: "Loutre",
		col_out: [][2]int{{0, 6}},
	},
	{
		search:  SearchInfo{Pattern: "Canapé", Granularity: LineGranularity},
		text:    "canapé",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "abc", Granularity: LineGranularity},
		text:    "loutre",
		exp_out: "",
	},
}

func TestSelectWholeLine(t *testing.T) {
	for _, st := range testsWholeLine {
		t.Run(
			fmt.Sprintf("find %#v in %#v", st.search, st.text),
			st.runST())
	}
}

var testsWords = []Subtest{
	{
		search:  SearchInfo{Pattern: "lo", Granularity: WordGranularity},
		text:    "loutre",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "Loutre", Granularity: WordGranularity},
		text:    "Loutre",
		exp_out: "Loutre",
		col_out: [][2]int{{0, 6}},
	},
	{
		search:  SearchInfo{Pattern: "Canapé", Granularity: WordGranularity},
		text:    "canapé",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "abc", Granularity: WordGranularity},
		text:    "loutre",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "fox", Granularity: WordGranularity},
		text:    "A fox called foxy",
		exp_out: "A fox called foxy",
		col_out: [][2]int{{2, 5}},
	},
	{
		search:  SearchInfo{Pattern: "user_1", Granularity: WordGranularity},
		text:    "user1 user10 user_10 user_1 user_1_1",
		exp_out: "user1 user10 user_10 user_1 user_1_1",
		col_out: [][2]int{{21, 27}},
	},
}

func TestSelectWords(t *testing.T) {

	for _, st := range testsWords {
		t.Run(
			fmt.Sprintf("find %#v in %#v", st.search, st.text),
			st.runST())
	}
}

var testsInsensitiveWord = []Subtest{
	{
		search:  SearchInfo{Pattern: "lo", Granularity: WordGranularity, CaseInsensitive: true},
		text:    "loutre",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "Loutre", Granularity: WordGranularity, CaseInsensitive: true},
		text:    "Loutre",
		exp_out: "Loutre",
		col_out: [][2]int{{0, 6}},
	},
	{
		search:  SearchInfo{Pattern: "Canapé", Granularity: WordGranularity, CaseInsensitive: true},
		text:    "canapé",
		exp_out: "canapé",
		col_out: [][2]int{{0, 7}},
	},
	{
		search:  SearchInfo{Pattern: "abc", Granularity: WordGranularity, CaseInsensitive: true},
		text:    "loutre",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "Fox", Granularity: WordGranularity, CaseInsensitive: true},
		text:    "A fox called Foxy",
		exp_out: "A fox called Foxy",
		col_out: [][2]int{{2, 5}},
	},
	{
		search:  SearchInfo{Pattern: "user_1", Granularity: WordGranularity, CaseInsensitive: true},
		text:    "user1 user10 user_10 user_1 user_1_1",
		exp_out: "user1 user10 user_10 user_1 user_1_1",
		col_out: [][2]int{{21, 27}},
	},
}

func TestSelectInsensitiveWords(t *testing.T) {

	for _, st := range testsInsensitiveWord {
		t.Run(
			fmt.Sprintf("find %#v in %#v", st.search, st.text),
			st.runST())
	}
}

var testsRegexpKeyword = []Subtest{
	{
		search:  SearchInfo{Pattern: `l\|t`},
		text:    "loutre",
		exp_out: "loutre",
		col_out: [][2]int{{0, 1}, {3, 4}},
	},
	{
		search:  SearchInfo{Pattern: `a\(b\|c\)d`},
		text:    "a abd acd abc",
		exp_out: "a abd acd abc",
		col_out: [][2]int{{2, 5}, {6, 9}},
	},
	{
		search:  SearchInfo{Pattern: `é\+`},
		text:    "éé èèè é ",
		exp_out: "éé èèè é ",
		col_out: [][2]int{{0, 4}, {12, 14}},
	},
	{
		search:  SearchInfo{Pattern: `ds\?`},
		text:    "adventure and making new friends.",
		exp_out: "adventure and making new friends.",
		col_out: [][2]int{{1, 2}, {12, 13}, {30, 32}},
	},
	{
		search:  SearchInfo{Pattern: `ds\+`},
		text:    "adventure and making new friends.",
		exp_out: "adventure and making new friends.",
		col_out: [][2]int{{30, 32}},
	},
	{
		search:  SearchInfo{Pattern: `ab*`},
		text:    "a ab abbb baa",
		exp_out: "a ab abbb baa",
		col_out: [][2]int{{0, 1}, {2, 4}, {5, 9}, {11, 12}, {12, 13}},
	},
	{
		search:  SearchInfo{Pattern: `\(D\|d\)`},
		text:    "Dis-donc tu es bien dodu",
		exp_out: "Dis-donc tu es bien dodu",
		col_out: [][2]int{{0, 1}, {4, 5}, {20, 21}, {22, 23}},
	},
	{
		search:  SearchInfo{Pattern: "."},
		text:    "ab+",
		exp_out: "ab+",
		col_out: [][2]int{{0, 1}, {1, 2}, {2, 3}},
	},
	{
		search:  SearchInfo{Pattern: ".*"},
		text:    "A fox called Foxy.",
		exp_out: "A fox called Foxy.",
		col_out: [][2]int{{0, 18}},
	},
	{
		search:  SearchInfo{Pattern: `d\{2\}`},
		text:    "hidden door",
		exp_out: "hidden door",
		col_out: [][2]int{{2, 4}},
	},
	{
		search:  SearchInfo{Pattern: `[abc]`},
		text:    "but also lasting bonds",
		exp_out: "but also lasting bonds",
		col_out: [][2]int{{0, 1}, {4, 5}, {10, 11}, {17, 18}},
	},
	{
		search:  SearchInfo{Pattern: `[^abc]`},
		text:    "but also",
		exp_out: "but also",
		col_out: [][2]int{{1, 2}, {2, 3}, {3, 4}, {5, 6}, {6, 7}, {7, 8}},
	},
	{
		search:  SearchInfo{Pattern: `[[:blank:]]`},
		text:    `Title: "Fantastic Fox Finds Friends"`,
		exp_out: `Title: "Fantastic Fox Finds Friends"`,
		col_out: [][2]int{{6, 7}, {17, 18}, {21, 22}, {27, 28}},
	},
	{
		search:  SearchInfo{Pattern: `[[:punct:]]`},
		text:    `Title: "Fantastic Fox Finds Friends"`,
		exp_out: `Title: "Fantastic Fox Finds Friends"`,
		col_out: [][2]int{{5, 6}, {7, 8}, {35, 36}},
	},
	{
		search:  SearchInfo{Pattern: `\"`},
		text:    `Title: "Fantastic Fox Finds Friends"`,
		exp_out: `Title: "Fantastic Fox Finds Friends"`,
		col_out: [][2]int{{7, 8}, {35, 36}},
	},
	{
		search:  SearchInfo{Pattern: `^A`},
		text:    `A Big Ananas`,
		exp_out: `A Big Ananas`,
		col_out: [][2]int{{0, 1}},
	},
	{
		search:  SearchInfo{Pattern: `s$`},
		text:    `saucisses`,
		exp_out: `saucisses`,
		col_out: [][2]int{{8, 9}},
	},
}

func TestSelectRegExpKeyword(t *testing.T) {

	for _, st := range testsRegexpKeyword {
		t.Run(
			fmt.Sprintf("find %#v in %#v", st.search, st.text),
			st.runST())
	}
}

func TestRegExpErrorHandling(t *testing.T) {
	var subtests = []struct {
		search SearchInfo
	}{
		{
			search: SearchInfo{Pattern: `\(D\|d`},
		},
	}

	for _, subtest := range subtests {
		var _, err_out = NewLineSelector(subtest.search)

		// we expect err_out != nil
		if err_out == nil {
			t.Errorf("wanted %#v (\"%v\"), got %#v,",
				nil, subtest.search, err_out)
		}
	}
}

func TestRegExpNotSupported(t *testing.T) {
	var subtests = []struct {
		search  SearchInfo
		text    string
		exp_out string
		col_out [][2]int
	}{
		{
			search:  SearchInfo{Pattern: `\<au\>`},
			text:    `saucisse au poulet`,
			exp_out: `saucisse au poulet`,
			col_out: [][2]int{{8, 9}},
		},
		{
			search:  SearchInfo{Pattern: `|`},
			text:    "loutre",
			exp_out: "",
			col_out: [][2]int{},
		},
	}
	search := SearchInfo{}

	for _, subtest := range subtests {
		var ls, err_out = NewLineSelector(search)

		if err_out != nil {
			t.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.search, err_out)
		}

		out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) {
			t.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v",
				subtest.exp_out, subtest.col_out,
				subtest.search, subtest.text,
				out.line, out.keywordRanges)
		}
	}
}

var testsOnlyMatching = []Subtest{

	{
		search:  SearchInfo{Pattern: "lo", OnlyMatching: true},
		text:    "loutre",
		exp_out: "lo",
		col_out: [][2]int{{0, 2}},
	},
	{
		search:  SearchInfo{Pattern: "Loutre", OnlyMatching: true},
		text:    "Loutre",
		exp_out: "Loutre",
		col_out: [][2]int{{0, 6}},
	},
	{
		search:  SearchInfo{Pattern: "Canapé", OnlyMatching: true},
		text:    "canapé",
		exp_out: "",
	},
	{
		search:  SearchInfo{Pattern: "a", OnlyMatching: true},
		text:    "canapé",
		exp_out: "a\na",
		col_out: [][2]int{{0, 1}, {2, 3}},
	},
	{
		search:  SearchInfo{Pattern: "abc", OnlyMatching: true},
		text:    "ABCabcABCabcABCàbç",
		exp_out: "abc\nabc",
		col_out: [][2]int{{0, 3}, {4, 7}},
	},
	{
		search:  SearchInfo{Pattern: "abc", OnlyMatching: true},
		text:    "loutre",
		exp_out: "",
	},
}

func TestSelectOnlyMatching(t *testing.T) {

	for _, st := range testsOnlyMatching {
		t.Run(
			fmt.Sprintf("find %#v in %#v", st.search, st.text),
			st.runST())
	}
}
