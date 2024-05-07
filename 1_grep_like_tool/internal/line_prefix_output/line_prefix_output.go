package line_prefix_output

import "main/internal/colors"

type LinePrefixRequest struct {
	WithFilename bool
}

type LinePrefixSelector struct {
	Lpr *LinePrefixRequest
}

func (lps LinePrefixSelector) GetPrefix(filename string) (ret string) {
	if lps.Lpr.WithFilename {
		ret = colors.Color_magenta(filename) + colors.Color_cyan(":")
	}
	return
}
