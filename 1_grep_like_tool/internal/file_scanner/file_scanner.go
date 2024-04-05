package grep_file_scanner

import (
	"bufio"
	"fmt"
	"io/fs"
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

func (fileScanner FileScanner) GoThroughFiles() (string, error) {

	var output string

	if !fileScanner.Recursive {
		for _, filename := range fileScanner.Paths {
			output += fileScanner.processFile(filename)
		}
	} else {
		for _, filename := range fileScanner.Paths {

			visitor := func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() {
					output += fileScanner.processFile(path)
				}
				return nil
			}

			err := filepath.WalkDir(filename, visitor)

			if err != nil {
				err = fmt.Errorf("error while walking dir : %w", err)
				return output, err
			}
		}
	}

	return output, nil
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
			} // else {
			// 	fmt.Fprintln(os.Stderr, "invalid line, not utf-8")
			// }

			// TODO : decide what to do with this
		}
	}
	file.Close()
	output += fs.Finder.OutputOnWholeFile(filename)
	return
}
