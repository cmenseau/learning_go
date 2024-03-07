package grep_line_select

import (
	grep_colors "main/internal/colors"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type SearchInfo struct {
	CaseInsensitive  bool
	InvertMatching   bool
	MatchGranularity string // default (""), word, line
}

func GetOutputLine(keyword string, line string, search SearchInfo) string {
	out_line, indexes_to_highlight := lineSelectorPipeline(keyword, line, search)
	return colorResults(out_line, indexes_to_highlight)
}

// returns string : output line
// returns [][2]int : indexes_to_highlight ([]byte indexes)
func lineSelectorPipeline(keyword string, line string, search SearchInfo) (string, [][2]int) {

	var indexes = getMatchingPatternIndexes(keyword, line)

	if !search.CaseInsensitive {
		indexes = applyCaseSensitiveSelection(keyword, line, indexes)
	}

	indexes = applyMatchGranularity(keyword, line, indexes, search.MatchGranularity)

	return applyInvertMatching(line, indexes, search.InvertMatching)
}

func colorResults(line string, indexes_to_highlight [][2]int) string {

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
		line_output += line[prev_end:start] + grep_colors.Color_red(line[start:end])
		prev_end = end
	}

	return line_output + line[prev_end:]
}

func applyInvertMatching(
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

func applyMatchGranularity(keyword string,
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

func applyCaseSensitiveSelection(keyword string, line string, indexes [][2]int) [][2]int {
	case_sensitive_indexes := [][2]int{}

	keyword_escape_removed := turnBreSyntaxInGoSyntax(keyword)
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

func getMatchingPatternIndexes(keyword string, line string) [][2]int {
	var indexes [][2]int
	// TODO turn this into [][]int like in regexp package ??

	keyword_escape_removed := turnBreSyntaxInGoSyntax(keyword)

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

func turnBreSyntaxInGoSyntax(keyword string) string {

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
