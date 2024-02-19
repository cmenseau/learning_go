package search_replace

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Perform() string {
	if len(os.Args[1:]) < 2 {
		fmt.Println("Parameters are missing")
	} else {
		old, new := os.Args[1], os.Args[2]
		scan := bufio.NewScanner(os.Stdin)
		replaced_string := ""

		var new_cap string = string(unicode.ToUpper([]rune(new)[0]))
		new_cap += string([]rune(new)[1:])
		var new_low string = string(unicode.ToLower([]rune(new)[0]))
		new_low += string([]rune(new)[1:])

		for scan.Scan() {
			var words []string = strings.Split(scan.Text(), " ")
			for i := range words {
				if strings.EqualFold(words[i], old) {
					// if strings.ToLower(words[i]) == strings.ToLower(old) {
					if unicode.IsUpper([]rune(words[i])[0]) {
						words[i] = new_cap
					} else {
						words[i] = new_low
					}
				}
			}
			replaced_string += strings.Join(words, " ") + "\n"
		}

		return replaced_string
	}
	return ""
}
