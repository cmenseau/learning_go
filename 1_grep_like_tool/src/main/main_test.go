package main

import (
	"testing"
)

func TestSelectHightlight(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
	}{
		{
			keyword: "lo",
			text:    "loutre",
			exp_out: color_red("lo") + "utre\n",
		},
		{
			keyword: "lo",
			text:    "Loutre",
			exp_out: "",
		},
		{
			keyword: "é",
			text:    "canapé",
			exp_out: "canap" + color_red("é") + "\n",
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "",
		},
	}

	for _, subtest := range subtests {
		var out = select_highlight(subtest.keyword, subtest.text)

		if out != subtest.exp_out {
			test.Errorf("wanted %#v (\"%v\" in : %v), got %#v",
				subtest.exp_out, subtest.keyword, subtest.text, out)
		}
	}
}

func TestSelectHightlightInsensitive(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
	}{
		{
			keyword: "lo",
			text:    "loutre",
			exp_out: color_red("lo") + "utre\n",
		},
		{
			keyword: "lo",
			text:    "Loutre",
			exp_out: color_red("Lo") + "utre\n",
		},
		{
			keyword: "é",
			text:    "canapé",
			exp_out: "canap" + color_red("é") + "\n",
		},
		{
			keyword: "abc",
			text:    "loutre",
			exp_out: "",
		},
	}

	for _, subtest := range subtests {
		var out = select_highlight_cis(subtest.keyword, subtest.text)

		if out != subtest.exp_out {
			test.Errorf("wanted %#v (\"%v\" in : %v), got %#v",
				subtest.exp_out, subtest.keyword, subtest.text, out)
		}
	}
}

func TestSelectReverse(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
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

	for _, subtest := range subtests {
		var out = select_reverse(subtest.keyword, subtest.text)

		if out != subtest.exp_out {
			test.Errorf("wanted %#v (\"%v\" in : %v), got %#v",
				subtest.exp_out, subtest.keyword, subtest.text, out)
		}
	}
}

func TestSelectWholeLine(test *testing.T) {
	var subtests = []struct {
		keyword string
		text    string
		exp_out string
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

	for _, subtest := range subtests {
		var out = select_whole_line(subtest.keyword, subtest.text)

		if out != subtest.exp_out {
			test.Errorf("wanted %#v (\"%v\" in : %v), got %#v",
				subtest.exp_out, subtest.keyword, subtest.text, out)
		}
	}
}
