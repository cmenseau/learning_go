package grep_line_prefix_control

import grep_colors "main/internal/colors"

type LinePrefixRequest struct {
	WithFilename bool
}

func (lpr LinePrefixRequest) GetPrefix(filename string) (ret string) {
	if lpr.WithFilename {
		ret = grep_colors.Color_magenta(filename) + grep_colors.Color_cyan(":")
	}
	return
}
