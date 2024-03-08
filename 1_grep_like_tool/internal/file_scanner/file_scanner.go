package grep_file_scanner

import (
	"bufio"
	"fmt"
	"os"
)

func GoThroughFiles(
	files []string,
	actionOnLine func(string, string) string,
	actionOnWholeFile func(string) string) string {

	var output string

	for _, filename := range files {

		file, err := os.Open(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open file %s because of %s", filename, err)
		} else {

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()

				output += actionOnLine(line, filename)
			}
		}
		file.Close()
		output += actionOnWholeFile(filename)
	}
	return output
}
