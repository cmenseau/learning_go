package core

import (
	"main/internal/file_output"
	"main/internal/line_output"
	"main/internal/line_prefix_output"
)

type Request struct {
	Paths      []string
	Recursive  bool
	Search     line_output.SearchInfo
	FileOutput file_output.FileOutputRequest
	LinePrefix line_prefix_output.LinePrefixRequest
}

func (r Request) IsRecursive() bool {
	return r.Recursive
}

func (r Request) GetPaths() []string {
	return r.Paths
}
