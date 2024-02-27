package main

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type search_info struct {
	case_insensitive  bool
	invert_matching   bool
	match_granularity string // default (""), word, line
}

func get_output_line(keyword string, line string, search search_info) string {
	out_line, indexes_to_highlight := line_selector_pipeline(keyword, line, search)
	return color_results(out_line, indexes_to_highlight)
}

// returns string : output line
// returns [][2]int : indexes_to_highlight ([]byte indexes)
func line_selector_pipeline(keyword string, line string, search search_info) (string, [][2]int) {

	var indexes = get_matching_pattern_indexes(keyword, line)

	if !search.case_insensitive {
		indexes = apply_case_sensitive_selection(keyword, line, indexes)
	}

	indexes = apply_match_granularity(keyword, line, indexes, search.match_granularity)

	return apply_invert_matching(line, indexes, search.invert_matching)
}

func color_results(line string, indexes_to_highlight [][2]int) string {

	if len(line) == 0 {
		return ""
	}
	if len(indexes_to_highlight) == 0 {
		// returns full line for empty indexes (invert_matching case) without color
		return line
	}

	var line_output = ""
	prev_end := 0

	for _, substr := range indexes_to_highlight {
		start, end := substr[0], substr[1]
		line_output += line[prev_end:start] + color_red(line[start:end])
		prev_end = end
	}

	return line_output + line[prev_end:]
}

func apply_invert_matching(
	line string,
	indexes [][2]int,
	invert_matching bool) (string, [][2]int) {

	var output = ""
	if invert_matching {
		// empty matches : line should be displayed, no highlight
		// matches found : line should not be displayed
		if len(indexes) != 0 {
			output = ""
			indexes = [][2]int{}
		} else {
			output = line
		}
	} else if len(indexes) > 0 {
		// at least 1 thing found : line should be displayed
		output = line
	}
	return output, indexes
}

func apply_match_granularity(keyword string,
	line string,
	indexes [][2]int,
	match_type string) [][2]int {

	seq_match_indexes := [][2]int{}

	if len(indexes) == 0 {
		return seq_match_indexes
	}

	switch match_type {
	case "":
		seq_match_indexes = indexes
	case "line":
		// indexes : [0]

		// if keyword is found more than 1 time, it can't be the full line
		// keyword must match full line (case insensitive)
		if len(indexes) == 1 && indexes[0][0] == 0 && strings.EqualFold(line, keyword) {
			seq_match_indexes = append(seq_match_indexes, indexes...)
		}
	case "word":

		for _, idx := range indexes {
			start := idx[0]
			stop := idx[1]
			var (
				has_space_before = false
				has_space_after  = false
			)
			if start == 0 {
				// front of the line
				has_space_before = true
			} else {
				previous_rune, _ := utf8.DecodeLastRuneInString(line[:start])
				if unicode.IsSpace(previous_rune) {
					has_space_before = true
				}
			}
			if stop == len(line) {
				// end of the line
				has_space_after = true
			} else {
				next_rune, _ := utf8.DecodeRuneInString(line[stop:]) // will decode 1st rune

				if unicode.IsSpace(next_rune) {
					has_space_after = true
				}
			}
			if has_space_before && has_space_after {
				seq_match_indexes = append(seq_match_indexes, idx)
			}
		}
	}

	return seq_match_indexes
}

func apply_case_sensitive_selection(keyword string, line string, indexes [][2]int) [][2]int {
	case_sensitive_indexes := [][2]int{}

	keyword_escape_removed := turn_bre_syntax_in_go_syntax(keyword)
	re, _ := regexp.Compile(keyword_escape_removed)
	// TODO err management

	for _, idx := range indexes {
		start := idx[0]
		stop := idx[1]
		if re.MatchString(line[start:stop]) {
			case_sensitive_indexes = append(case_sensitive_indexes, idx)
		}

		//TODO what about cases when keyword contains regular expression using surrounding context
	}
	return case_sensitive_indexes
}

func get_matching_pattern_indexes(keyword string, line string) [][2]int {
	var indexes [][2]int
	// TODO turn this into [][]int like in regexp package ??

	keyword_escape_removed := turn_bre_syntax_in_go_syntax(keyword)

	var lower_keyword = strings.ToLower(keyword_escape_removed)
	var lower_line = strings.ToLower(line)

	curr_line := lower_line

	re, err := regexp.Compile(lower_keyword)
	if err == nil {
		for _, idx_pair := range re.FindAllStringIndex(curr_line, -1) {
			indexes = append(indexes, [2]int{idx_pair[0], idx_pair[1]})
		}
	} else {
		// TODO
	}
	return indexes
}

func turn_bre_syntax_in_go_syntax(keyword string) string {

	// replace \| \? \+ \( \) \{ \}
	// by      |   ?  +  (  )  {  }

	re_escape := regexp.MustCompile(`\\\||\\\?|\\\+|\\\(|\\\)|\\\{|\\\}`)
	keyword_escape_removed := re_escape.ReplaceAllStringFunc(keyword, func(s string) string {
		switch s {
		case `\|`:
			return `|`
		case `\?`:
			return `?`
		case `\+`:
			return `+`
		case `\(`:
			return `(`
		case `\)`:
			return `)`
		case `\{`:
			return `{`
		case `\}`:
			return `}`
		default:
			return s
		}
	})
	return keyword_escape_removed
}
