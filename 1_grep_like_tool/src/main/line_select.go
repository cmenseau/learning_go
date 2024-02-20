package main

import (
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

func line_selector_pipeline(keyword string, line string, search search_info) (string, [][2]int) {
	var output = ""

	var indexes = get_matching_pattern_indexes(keyword, line)

	if !search.case_insensitive {
		indexes = apply_case_sensitive_selection(keyword, line, indexes)
	}

	indexes = apply_match_granularity(keyword, line, indexes, search.match_granularity)

	if search.invert_matching {
		// empty matches : line should be displayed, no highlight
		// matches found : line should not be displayed
		if len(indexes) != 0 {
			output = ""
			indexes = [][2]int{}
		} else {
			output = line + "\n"
		}
	} else {
		if len(indexes) > 0 {
			output = line + "\n"
		}
	}
	return output, indexes
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

	for _, idx := range indexes {
		start := idx[0]
		stop := idx[1]
		if line[start:stop] == keyword {
			case_sensitive_indexes = append(case_sensitive_indexes, idx)
		}
	}
	return case_sensitive_indexes
}

func get_matching_pattern_indexes(keyword string, line string) [][2]int {
	var indexes [][2]int

	var lower_keyword = strings.ToLower(keyword)
	var lower_line = strings.ToLower(line)

	curr_line := lower_line
	curr_idx := strings.Index(curr_line, lower_keyword)

	// while we find the keyword pattern, we add its index in the indexes list
	for curr_idx != -1 {
		indexes = append(indexes, [2]int{curr_idx, curr_idx + len(lower_keyword)})
		curr_line = lower_line[curr_idx+len(lower_keyword):]
		next_keyword_idx := strings.Index(curr_line, lower_keyword)
		if next_keyword_idx == -1 {
			curr_idx = -1
		} else {
			curr_idx = next_keyword_idx + curr_idx + len(lower_keyword)
		}
	}
	return indexes
}
