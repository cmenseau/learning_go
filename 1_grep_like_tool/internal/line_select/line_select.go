package grep_line_select

import (
	"fmt"
	grep_colors "main/internal/colors"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type matchGranularity int

const (
	AllGranularity matchGranularity = iota
	WordGranularity
	LineGranularity
)

type SearchInfo struct {
	CaseInsensitive bool
	InvertMatching  bool
	Granularity     matchGranularity
	OnlyMatching    bool
}

// intermediate struct to store result line
type highlightedLine struct {
	line          string
	keywordRanges [][2]int
}

type LineSelector struct {
	Keyword            string
	Search             SearchInfo
	regexp             *regexp.Regexp
	regexCaseSensitive *regexp.Regexp
}

func NewLineSelector(keyword string, search SearchInfo) (LineSelector, error) {
	var ls LineSelector
	ls.Keyword = keyword
	ls.Search = search

	var keyword_escape_removed = turnBreSyntaxInGoSyntax(ls.Keyword)
	var lower_keyword = strings.ToLower(keyword_escape_removed)

	re, err := regexp.Compile(lower_keyword)

	if err != nil {
		return LineSelector{}, fmt.Errorf("using given regex (%s) : %w", ls.Keyword, err)
	}
	ls.regexp = re

	if !ls.Search.CaseInsensitive {
		re, err = regexp.Compile(keyword_escape_removed)

		if err != nil {
			return LineSelector{}, fmt.Errorf("using given regex (%s) : %w", ls.Keyword, err)
		}
		ls.regexCaseSensitive = re
	}

	return ls, nil
}

func (ls LineSelector) GetOutputLine(line string) (string, error) {
	out_line, err := ls.lineSelectorPipeline(line)
	return colorResults(out_line), err
}

func (ls LineSelector) lineSelectorPipeline(line string) (output highlightedLine, err error) {

	indexes, err := ls.getMatchingPatternIndexes(line)

	if err != nil {
		err = fmt.Errorf("can't get matching pattern : %w", err)
		return
	}

	if !ls.Search.CaseInsensitive {
		indexes, err = ls.applyCaseSensitiveSelection(line, indexes)

		if err != nil {
			err = fmt.Errorf("can't select case sensitively : %w", err)
			return
		}
	}

	indexes = ls.applyMatchGranularity(line, indexes)

	if len(indexes) > 0 {
		// at least 1 thing found : line should be displayed
		output.line = line
		output.keywordRanges = indexes
	}

	if ls.Search.InvertMatching {
		output_line, indexes := applyInvertMatching(line, indexes)
		output.line = output_line
		output.keywordRanges = indexes
	}

	if ls.Search.OnlyMatching {
		output_line, indexes := applyOnlyMatching(line, indexes)
		output.line = output_line
		output.keywordRanges = indexes
	}

	return output, nil
	// we need to return both line and indexes because line could
	// be selected without any highlighted part (indexes={})
}

func colorResults(lineToHighlight highlightedLine) string {

	if len(lineToHighlight.line) == 0 {
		return ""
	}
	if len(lineToHighlight.keywordRanges) == 0 {
		// returns full line for empty indexes (invert_matching case) without color
		return lineToHighlight.line
	}

	var line_output = ""
	prev_end := 0

	for _, substr := range lineToHighlight.keywordRanges {
		start, end := substr[0], substr[1]
		line_output += lineToHighlight.line[prev_end:start] + grep_colors.Color_red(lineToHighlight.line[start:end])
		prev_end = end
	}

	return line_output + lineToHighlight.line[prev_end:]
}

func applyInvertMatching(
	line string,
	indexes [][2]int) (string, [][2]int) {

	// empty matches : line should be displayed, no highlight
	// matches found : line should not be displayed (no highlight)
	var output = ""
	if len(indexes) == 0 {
		output = line
	}
	return output, nil
}

func applyOnlyMatching(
	line string,
	indexes_to_highlight [][2]int) (string, [][2]int) {

	// remove everything that's not highlighted
	var output = ""
	var new_indexes [][2]int

	for idx, substr := range indexes_to_highlight {
		line_idx := len(output)
		output += line[substr[0]:substr[1]]

		if idx != len(indexes_to_highlight)-1 {
			output += "\n"
		}
		new_indexes = append(new_indexes,
			[2]int{line_idx, line_idx + substr[1] - substr[0]})
	}

	return output, new_indexes
}

func (ls LineSelector) applyMatchGranularity(
	line string,
	indexes [][2]int) [][2]int {

	seq_match_indexes := [][2]int{}

	if len(indexes) == 0 {
		return seq_match_indexes
	}

	switch ls.Search.Granularity {
	case AllGranularity:
		seq_match_indexes = indexes
	case LineGranularity:
		// indexes : [0]

		// if keyword is found more than 1 time, it can't be the full line
		// keyword must match full line (case insensitive)
		if len(indexes) == 1 && indexes[0][0] == 0 && strings.EqualFold(line, ls.Keyword) {
			seq_match_indexes = append(seq_match_indexes, indexes...)
		}
	case WordGranularity:

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

func (ls LineSelector) applyCaseSensitiveSelection(line string, indexes [][2]int) (case_sensitive_indexes [][2]int, err error) {

	// TODO
	if ls.Search.CaseInsensitive {
		panic("programming error")
	}

	for _, idx := range indexes {
		start := idx[0]
		stop := idx[1]
		if ls.regexCaseSensitive.MatchString(line[start:stop]) {
			case_sensitive_indexes = append(case_sensitive_indexes, idx)
		}

		//TODO what about cases when keyword contains regular expression using surrounding context
	}
	return case_sensitive_indexes, nil
}

func (ls LineSelector) getMatchingPatternIndexes(line string) (indexes [][2]int, err error) {

	var lower_line = strings.ToLower(line)

	for _, idx_pair := range ls.regexp.FindAllStringIndex(lower_line, -1) {
		indexes = append(indexes, [2]int{idx_pair[0], idx_pair[1]})
	}
	return indexes, nil
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
