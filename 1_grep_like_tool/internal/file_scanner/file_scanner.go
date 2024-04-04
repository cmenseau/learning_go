package grep_file_scanner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"unicode/utf8"
)

type Finder interface {
	OutputOnLine(line string, filename string) string
	OutputOnWholeFile(filename string) string
}

type FileScanner struct {
	Finder    Finder
	Paths     []string // can be both file paths or dir paths
	Recursive bool
}

func (fs FileScanner) GoThroughFiles() string {

	var output string

	if !fs.Recursive {
		for _, filename := range fs.Paths {
			output += fs.processFile(filename)
		}
	} else {
		for _, filename := range fs.Paths {

			visitor := func(path string, fi os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !fi.IsDir() {
					output += fs.processFile(path)
				}
				return nil
			}

			filepath.Walk(filename, visitor)
		}
	}

	return output
}

func (fs FileScanner) processFile(filename string) (output string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open file %s because of %s", filename, err)
	} else {

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			if utf8.ValidString(line) {
				output += fs.Finder.OutputOnLine(line, filename)
			} else {
				fmt.Fprintln(os.Stderr, "invalid line, not utf-8")
			}
		}
	}
	file.Close()
	output += fs.Finder.OutputOnWholeFile(filename)
	return
}
