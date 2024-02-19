package hi

import (
	"strings"
)

func Say(names []string) string {
	if len(names) == 0 {
		return "Hello World!"
	} else {
		return "Hello " + strings.Join(names, ", ") + "!"
	}
}
