package main

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
			exp_out: "loutre\n",
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
			exp_out: "canapé\n",
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
			exp_out: "loutre à clou\n",
			col_out: [][2]int{{0, 2}, {11, 13}}, // à is 2 bytes
		},
	}
	search := search_info{}

	for _, subtest := range subtests {
		var str_out, idx_out = line_selector_pipeline(subtest.keyword, subtest.text, search)

		if str_out != subtest.exp_out || !slices.Equal(idx_out, subtest.col_out) {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				str_out, idx_out)
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
			exp_out: "loutre\n",
			col_out: [][2]int{{0, 2}},
		},
		{
			keyword: "lo",
			text:    "Loutre",
			exp_out: "Loutre\n",
			col_out: [][2]int{{0, 2}},
		},
		{
			keyword: "é",
			text:    "canapé",
			exp_out: "canapé\n",
			col_out: [][2]int{{5, 7}}, // é is 2 bytes
		},
		{
			keyword: "É",
			text:    "canapé",
			exp_out: "canapé\n",
			col_out: [][2]int{{5, 7}}, // é is 2 bytes
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "",
		},
	}

	search := search_info{case_insensitive: true}

	for _, subtest := range subtests {
		var str_out, idx_out = line_selector_pipeline(subtest.keyword, subtest.text, search)

		if str_out != subtest.exp_out || !slices.Equal(idx_out, subtest.col_out) {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				str_out, idx_out)
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
			exp_out: "Loutre\n",
		},
		{
			keyword: "é",
			text:    "canapé",
			exp_out: "",
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "loutre\n",
		},
	}

	search := search_info{invert_matching: true}

	for _, subtest := range subtests {
		var str_out, idx_out = line_selector_pipeline(subtest.keyword, subtest.text, search)

		if str_out != subtest.exp_out || !slices.Equal(idx_out, subtest.col_out) {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				str_out, idx_out)
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
			exp_out: "Loutre\n",
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
	search := search_info{match_granularity: "line"}

	for _, subtest := range subtests {
		var str_out, idx_out = line_selector_pipeline(subtest.keyword, subtest.text, search)

		if str_out != subtest.exp_out || !slices.Equal(idx_out, subtest.col_out) {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				str_out, idx_out)
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
			exp_out: "Loutre\n",
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
			exp_out: "A fox called foxy\n",
			col_out: [][2]int{{2, 5}},
		},
		{
			keyword: "user_1",
			text:    "user1 user10 user_10 user_1 user_1_1",
			exp_out: "user1 user10 user_10 user_1 user_1_1\n",
			col_out: [][2]int{{21, 27}},
		},
	}
	search := search_info{match_granularity: "word"}

	for _, subtest := range subtests {
		var str_out, idx_out = line_selector_pipeline(subtest.keyword, subtest.text, search)

		if str_out != subtest.exp_out || !slices.Equal(idx_out, subtest.col_out) {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				str_out, idx_out)
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
			exp_out: "Loutre\n",
			col_out: [][2]int{{0, 6}},
		},
		{
			keyword: "Canapé",
			text:    "canapé",
			exp_out: "canapé\n",
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
			exp_out: "A fox called Foxy\n",
			col_out: [][2]int{{2, 5}},
		},
		{
			keyword: "user_1",
			text:    "user1 user10 user_10 user_1 user_1_1",
			exp_out: "user1 user10 user_10 user_1 user_1_1\n",
			col_out: [][2]int{{21, 27}},
		},
	}
	search := search_info{match_granularity: "word"}
	search.case_insensitive = true

	for _, subtest := range subtests {
		var str_out, idx_out = line_selector_pipeline(subtest.keyword, subtest.text, search)

		if str_out != subtest.exp_out || !slices.Equal(idx_out, subtest.col_out) {
			test.Errorf("wanted %#v, %#v (\"%v\" in : %v), got %#v, %#v",
				subtest.exp_out, subtest.col_out,
				subtest.keyword, subtest.text,
				str_out, idx_out)
		}
	}
}
