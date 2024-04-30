package line_prefix_output

import "main/internal/colors"

type LinePrefixRequest struct {
	WithFilename bool
}

func (lpr LinePrefixRequest) GetPrefix(filename string) (ret string) {
	if lpr.WithFilename {
		ret = colors.Color_magenta(filename) + colors.Color_cyan(":")
	}
	return
}
