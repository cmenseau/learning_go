package grep_file_scanner

import (
	"bufio"
	"fmt"
	"os"
)

type Finder interface {
	OutputOnLine(line string, filename string) string
	OutputOnWholeFile(filename string) string
}

type FileScanner struct {
	Finder    Finder
	Filenames []string
}

func (fs FileScanner) GoThroughFiles() string {

	var output string

	for _, filename := range fs.Filenames {

		file, err := os.Open(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open file %s because of %s", filename, err)
		} else {

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()

				output += fs.Finder.OutputOnLine(line, filename)
			}
		}
		file.Close()
		output += fs.Finder.OutputOnWholeFile(filename)
	}
	return output
}
