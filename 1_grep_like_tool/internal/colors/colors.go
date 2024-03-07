package grep_colors

import "github.com/fatih/color"

/////// COLOR OUTPUT FUNCTIONS ///////////

func Color_red(str string) string {
	red := color.New(color.FgRed).SprintFunc()
	return red(str)
}

func Color_magenta(str string) string {
	magenta := color.New(color.FgMagenta).SprintFunc()
	return magenta(str)
}

func Color_cyan(str string) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	return cyan(str)
}
