package parser

import (
	"errors"
	"fmt"
	"main/internal/engine"
	"main/internal/line_output"
	"os"
	"strings"
)

type InputModel struct {
	Pattern string
	Paths   []string
	InputOptions
}

type InputOptions struct {
	Recursive         bool // -r
	IgnoreCase        bool // -i
	InvertMatch       bool // -v
	WordRegexp        bool // -w
	LineRegexp        bool // -x
	OnlyMatching      bool // -o
	Count             bool // -c
	FilesWithoutMatch bool // -L
	FilesWithMatches  bool // -l
	WithFilename      bool // -H
}

// parse args and returns GrepRequest
func ParseArgs(args []string) (req engine.Request, err error) {

	if len(args) < 2 {
		err = fmt.Errorf("expecting more than 2 args, got %v", os.Args[1:])
		return
	}

	var options []string
	var pattern_idx int
	var first_filename_idx int
	var pattern_found, filename_found bool

	for i, val := range args {
		if strings.HasPrefix(val, "-") {
			var chars = strings.Split(strings.Replace(val, "-", "", -1), "")
			options = append(options, chars...)
		} else {
			// not an option : either the pattern if it wasn't yet found,
			// otherwise the first filepath
			if !pattern_found {
				pattern_idx = i
				pattern_found = true
			} else {
				first_filename_idx = i
				filename_found = true
				break
			}
		}
	}

	if !pattern_found || !filename_found {
		err = errors.New("expecting arguments : [options...] pattern file... ")
		return
	}

	model := InputModel{}

	model.Pattern = args[pattern_idx]
	model.Paths = args[first_filename_idx:]

	for _, opt := range options {
		switch opt {
		case "i":
			model.IgnoreCase = true
		case "w":
			model.WordRegexp = true
		case "x":
			model.LineRegexp = true
		case "v":
			model.InvertMatch = true
		case "r":
			model.Recursive = true
		case "o":
			model.OnlyMatching = true
		case "c":
			model.Count = true
		case "L":
			model.FilesWithoutMatch = true
		case "l":
			model.FilesWithMatches = true
		case "H":
			model.WithFilename = true
		default:
			err = fmt.Errorf("unsupported option %s", opt)
			return
		}
	}

	req = *modelToRequest(model)

	return
}

func modelToRequest(input InputModel) *engine.Request {
	r := engine.Request{}
	r.Paths = input.Paths
	r.Search.Pattern = input.Pattern

	if len(input.Paths) > 1 || input.WithFilename {
		r.LinePrefix.WithFilename = true
	}

	// -r with only 1 filename : r.LinePrefix.WithFilename => false
	// -r with 1+ dir : r.LinePrefix.WithFilename => true
	if input.Recursive {
		r.Recursive = true

		if fi, err := os.Stat(r.Paths[0]); len(r.Paths) == 1 && err == nil && !fi.IsDir() {
			r.LinePrefix.WithFilename = false
		} else {
			r.LinePrefix.WithFilename = true
		}
	}

	if input.IgnoreCase {
		r.Search.CaseInsensitive = true
	}
	if input.InvertMatch {
		r.Search.InvertMatching = true
	}
	if input.OnlyMatching {
		r.Search.OnlyMatching = true
	}

	// if line regexp and word regexp specified, line regexp takes over
	if input.WordRegexp {
		r.Search.Granularity = line_output.WordGranularity
	}
	if input.LineRegexp {
		r.Search.Granularity = line_output.LineGranularity
	}

	// only one the the 3 option can be used,
	// with priority filesWithout over filesWith over CountLines
	if input.FilesWithoutMatch {
		r.FileOutput.FilesWithoutMatch = true
		r.LinePrefix.WithFilename = false
	} else if input.FilesWithMatches {
		r.FileOutput.FilesWithMatch = true
		r.LinePrefix.WithFilename = false
	} else if input.Count {
		r.FileOutput.CountLines = true
	}

	return &r
}
