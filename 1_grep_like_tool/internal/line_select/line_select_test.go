package grep_line_select

import (
	"slices"
	"testing"
)

func TestSelectHightlight(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
		col_out [][2]int
	}{
		{
			keyword: "lo",
			text:    "loutre",
			exp_out: "loutre",
			col_out: [][2]int{{0, 2}},
		},
		{
			keyword: "lo",
			text:    "Loutre",
			exp_out: "",
		},
		{
			keyword: "é",
			text:    "canapé",
			exp_out: "canapé",
			col_out: [][2]int{{5, 7}}, // é is 2 bytes
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "",
		},
		{
			keyword: "lo",
			text:    "loutre à clou",
			exp_out: "loutre à clou",
			col_out: [][2]int{{0, 2}, {11, 13}}, // à is 2 bytes
		},
	}
	search := SearchInfo{}

	for _, subtest := range subtests {
		var ls, err_out = NewLineSelector(subtest.keyword, search)

		if err_out != nil {
			test.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.keyword, err_out)
		}

		out, err_out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) || err_out != nil {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v, %v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				out.line, out.keywordRanges, err_out)
		}
	}
}

func TestSelectHightlightInsensitive(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
		col_out [][2]int
	}{
		{
			keyword: "lo",
			text:    "loutre",
			exp_out: "loutre",
			col_out: [][2]int{{0, 2}},
		},
		{
			keyword: "lo",
			text:    "Loutre",
			exp_out: "Loutre",
			col_out: [][2]int{{0, 2}},
		},
		{
			keyword: "é",
			text:    "canapé",
			exp_out: "canapé",
			col_out: [][2]int{{5, 7}}, // é is 2 bytes
		},
		{
			keyword: "É",
			text:    "canapé",
			exp_out: "canapé",
			col_out: [][2]int{{5, 7}}, // é is 2 bytes
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "",
		},
	}

	search := SearchInfo{CaseInsensitive: true}

	for _, subtest := range subtests {
		var ls, err_out = NewLineSelector(subtest.keyword, search)

		if err_out != nil {
			test.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.keyword, err_out)
		}

		out, err_out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) || err_out != nil {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v, %v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				out.line, out.keywordRanges, err_out)
		}
	}
}

func TestSelectReverse(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
		col_out [][2]int
	}{
		{
			keyword: "lo",
			text:    "loutre",
			exp_out: "",
		},
		{
			keyword: "lo",
			text:    "Loutre",
			exp_out: "Loutre",
		},
		{
			keyword: "é",
			text:    "canapé",
			exp_out: "",
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "loutre",
		},
	}

	search := SearchInfo{InvertMatching: true}

	for _, subtest := range subtests {
		var ls, err_out = NewLineSelector(subtest.keyword, search)

		if err_out != nil {
			test.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.keyword, err_out)
		}

		out, err_out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) || err_out != nil {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v, %v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				out.line, out.keywordRanges, err_out)
		}
	}
}

func TestSelectWholeLine(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
		col_out [][2]int
	}{
		{
			keyword: "lo",
			text:    "loutre",
			exp_out: "",
		},
		{
			keyword: "Loutre",
			text:    "Loutre",
			exp_out: "Loutre",
			col_out: [][2]int{{0, 6}},
		},
		{
			keyword: "Canapé",
			text:    "canapé",
			exp_out: "",
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "",
		},
	}
	search := SearchInfo{Granularity: LineGranularity}

	for _, subtest := range subtests {
		var ls, err_out = NewLineSelector(subtest.keyword, search)

		if err_out != nil {
			test.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.keyword, err_out)
		}

		out, err_out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) || err_out != nil {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v, %v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				out.line, out.keywordRanges, err_out)
		}
	}
}

func TestSelectWords(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
		col_out [][2]int
	}{
		{
			keyword: "lo",
			text:    "loutre",
			exp_out: "",
		},
		{
			keyword: "Loutre",
			text:    "Loutre",
			exp_out: "Loutre",
			col_out: [][2]int{{0, 6}},
		},
		{
			keyword: "Canapé",
			text:    "canapé",
			exp_out: "",
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "",
		},
		{
			keyword: "fox",
			text:    "A fox called foxy",
			exp_out: "A fox called foxy",
			col_out: [][2]int{{2, 5}},
		},
		{
			keyword: "user_1",
			text:    "user1 user10 user_10 user_1 user_1_1",
			exp_out: "user1 user10 user_10 user_1 user_1_1",
			col_out: [][2]int{{21, 27}},
		},
	}
	search := SearchInfo{Granularity: WordGranularity}

	for _, subtest := range subtests {
		var ls, err_out = NewLineSelector(subtest.keyword, search)

		if err_out != nil {
			test.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.keyword, err_out)
		}

		out, err_out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) || err_out != nil {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v, %v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				out.line, out.keywordRanges, err_out)
		}
	}
}

func TestSelectInsensitiveWords(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
		col_out [][2]int
	}{
		{
			keyword: "lo",
			text:    "loutre",
			exp_out: "",
		},
		{
			keyword: "Loutre",
			text:    "Loutre",
			exp_out: "Loutre",
			col_out: [][2]int{{0, 6}},
		},
		{
			keyword: "Canapé",
			text:    "canapé",
			exp_out: "canapé",
			col_out: [][2]int{{0, 7}},
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "",
		},
		{
			keyword: "Fox",
			text:    "A fox called Foxy",
			exp_out: "A fox called Foxy",
			col_out: [][2]int{{2, 5}},
		},
		{
			keyword: "user_1",
			text:    "user1 user10 user_10 user_1 user_1_1",
			exp_out: "user1 user10 user_10 user_1 user_1_1",
			col_out: [][2]int{{21, 27}},
		},
	}
	search := SearchInfo{Granularity: WordGranularity}
	search.CaseInsensitive = true

	for _, subtest := range subtests {
		var ls, err_out = NewLineSelector(subtest.keyword, search)

		if err_out != nil {
			test.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.keyword, err_out)
		}

		out, err_out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) || err_out != nil {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v, %v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				out.line, out.keywordRanges, err_out)
		}
	}
}

func TestSelectRegExpKeyword(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
		col_out [][2]int
	}{
		{
			keyword: `l\|t`,
			text:    "loutre",
			exp_out: "loutre",
			col_out: [][2]int{{0, 1}, {3, 4}},
		},
		{
			keyword: `a\(b\|c\)d`,
			text:    "a abd acd abc",
			exp_out: "a abd acd abc",
			col_out: [][2]int{{2, 5}, {6, 9}},
		},
		{
			keyword: `é\+`,
			text:    "éé èèè é ",
			exp_out: "éé èèè é ",
			col_out: [][2]int{{0, 4}, {12, 14}},
		},
		{
			keyword: `ds\?`,
			text:    "adventure and making new friends.",
			exp_out: "adventure and making new friends.",
			col_out: [][2]int{{1, 2}, {12, 13}, {30, 32}},
		},
		{
			keyword: `ds\+`,
			text:    "adventure and making new friends.",
			exp_out: "adventure and making new friends.",
			col_out: [][2]int{{30, 32}},
		},
		{
			keyword: `ab*`,
			text:    "a ab abbb baa",
			exp_out: "a ab abbb baa",
			col_out: [][2]int{{0, 1}, {2, 4}, {5, 9}, {11, 12}, {12, 13}},
		},
		{
			keyword: `\(D\|d\)`,
			text:    "Dis-donc tu es bien dodu",
			exp_out: "Dis-donc tu es bien dodu",
			col_out: [][2]int{{0, 1}, {4, 5}, {20, 21}, {22, 23}},
		},
		{
			keyword: ".",
			text:    "ab+",
			exp_out: "ab+",
			col_out: [][2]int{{0, 1}, {1, 2}, {2, 3}},
		},
		{
			keyword: ".*",
			text:    "A fox called Foxy.",
			exp_out: "A fox called Foxy.",
			col_out: [][2]int{{0, 18}},
		},
		{
			keyword: `d\{2\}`,
			text:    "hidden door",
			exp_out: "hidden door",
			col_out: [][2]int{{2, 4}},
		},
		{
			keyword: `[abc]`,
			text:    "but also lasting bonds",
			exp_out: "but also lasting bonds",
			col_out: [][2]int{{0, 1}, {4, 5}, {10, 11}, {17, 18}},
		},
		{
			keyword: `[^abc]`,
			text:    "but also",
			exp_out: "but also",
			col_out: [][2]int{{1, 2}, {2, 3}, {3, 4}, {5, 6}, {6, 7}, {7, 8}},
		},
		{
			keyword: `[[:blank:]]`,
			text:    `Title: "Fantastic Fox Finds Friends"`,
			exp_out: `Title: "Fantastic Fox Finds Friends"`,
			col_out: [][2]int{{6, 7}, {17, 18}, {21, 22}, {27, 28}},
		},
		{
			keyword: `[[:punct:]]`,
			text:    `Title: "Fantastic Fox Finds Friends"`,
			exp_out: `Title: "Fantastic Fox Finds Friends"`,
			col_out: [][2]int{{5, 6}, {7, 8}, {35, 36}},
		},
		{
			keyword: `\"`,
			text:    `Title: "Fantastic Fox Finds Friends"`,
			exp_out: `Title: "Fantastic Fox Finds Friends"`,
			col_out: [][2]int{{7, 8}, {35, 36}},
		},
		{
			keyword: `^A`,
			text:    `A Big Ananas`,
			exp_out: `A Big Ananas`,
			col_out: [][2]int{{0, 1}},
		},
		{
			keyword: `s$`,
			text:    `saucisses`,
			exp_out: `saucisses`,
			col_out: [][2]int{{8, 9}},
		},
	}
	search := SearchInfo{}

	for _, subtest := range subtests {
		var ls, err_out = NewLineSelector(subtest.keyword, search)

		if err_out != nil {
			test.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.keyword, err_out)
		}

		out, err_out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) || err_out != nil {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v, %v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				out.line, out.keywordRanges, err_out)
		}
	}
}

func TestRegExpErrorHandling(test *testing.T) {
	var subtests = []struct {
		keyword string
	}{
		{
			keyword: `\(D\|d`,
		},
	}
	search := SearchInfo{}

	for _, subtest := range subtests {
		var _, err_out = NewLineSelector(subtest.keyword, search)

		// we expect err_out != nil
		if err_out == nil {
			test.Errorf("wanted %#v (\"%v\"), got %#v,",
				nil, subtest.keyword, err_out)
		}
	}
}

func TestRegExpNotSupported(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
		col_out [][2]int
	}{
		{
			keyword: `\<au\>`,
			text:    `saucisse au poulet`,
			exp_out: `saucisse au poulet`,
			col_out: [][2]int{{8, 9}},
		},
		{
			keyword: `|`,
			text:    "loutre",
			exp_out: "",
			col_out: [][2]int{},
		},
	}
	search := SearchInfo{}

	for _, subtest := range subtests {
		var ls, err_out = NewLineSelector(subtest.keyword, search)

		if err_out != nil {
			test.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.keyword, err_out)
		}

		out, err_out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) || err_out != nil {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v, %v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				out.line, out.keywordRanges, err_out)
		}
	}
}

func TestSelectOnlyMatching(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
		col_out [][2]int
	}{
		{
			keyword: "lo",
			text:    "loutre",
			exp_out: "lo",
			col_out: [][2]int{{0, 2}},
		},
		{
			keyword: "Loutre",
			text:    "Loutre",
			exp_out: "Loutre",
			col_out: [][2]int{{0, 6}},
		},
		{
			keyword: "Canapé",
			text:    "canapé",
			exp_out: "",
		},
		{
			keyword: "a",
			text:    "canapé",
			exp_out: "a\na",
			col_out: [][2]int{{0, 1}, {2, 3}},
		},
		{
			keyword: "abc",
			text:    "ABCabcABCabcABCàbç",
			exp_out: "abc\nabc",
			col_out: [][2]int{{0, 3}, {4, 7}},
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "",
		},
	}

	search := SearchInfo{OnlyMatching: true}

	for _, subtest := range subtests {
		var ls, err_out = NewLineSelector(subtest.keyword, search)

		if err_out != nil {
			test.Errorf("wanted %#v (\"%v\"), got %#v",
				nil, subtest.keyword, err_out)
		}

		out, err_out := ls.lineSelectorPipeline(subtest.text)

		if out.line != subtest.exp_out || !slices.Equal(out.keywordRanges, subtest.col_out) || err_out != nil {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v, %v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				out.line, out.keywordRanges, err_out)
		}
	}
}
