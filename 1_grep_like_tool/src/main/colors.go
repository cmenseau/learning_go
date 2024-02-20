package main

import "github.com/fatih/color"

/////// COLOR OUTPUT FUNCTIONS ///////////

func color_red(str string) string {
	red := color.New(color.FgRed).SprintFunc()
	return red(str)
}

func color_magenta(str string) string {
	magenta := color.New(color.FgMagenta).SprintFunc()
	return magenta(str)
}

func color_cyan(str string) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	return cyan(str)
}
