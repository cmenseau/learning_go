package grep_output_control

import "fmt"

type FileOutputRequest struct {
	CountLines            bool
	FilesWithoutMatch     bool
	FilesWithMatch        bool
	filesWithoutMatchMap  map[string]bool // set of files without match
	filesWithMatchMap     map[string]bool // set of files with matches
	countMatchingLinesMap map[string]int
}

func (fos FileOutputRequest) SuppressNormalOutput() bool {
	return fos.CountLines ||
		fos.FilesWithoutMatch ||
		fos.FilesWithMatch
}

func (fos *FileOutputRequest) ProcessOutputLine(line string, filename string) {
	// store count or what's required to be able to print final output by file

	switch {
	case fos.FilesWithoutMatch:
		if fos.filesWithoutMatchMap == nil {
			fos.filesWithoutMatchMap = make(map[string]bool)
		}
		if line == "" {
			fos.filesWithoutMatchMap[filename] = true
		}
	case fos.FilesWithMatch:
		if fos.filesWithMatchMap == nil {
			fos.filesWithMatchMap = make(map[string]bool)
		}
		if line != "" {
			fos.filesWithMatchMap[filename] = true
		}
	case fos.CountLines:
		if fos.countMatchingLinesMap == nil {
			fos.countMatchingLinesMap = make(map[string]int)
		}
		if line != "" {
			fos.countMatchingLinesMap[filename] += 1
		}
	}
}

func (fos *FileOutputRequest) GetFinalOutputControl(filename string) string {
	switch {
	case fos.FilesWithoutMatch:
		if fos.filesWithoutMatchMap[filename] {
			return filename
		}
	case fos.FilesWithMatch:
		if fos.filesWithMatchMap[filename] {
			return filename
		}
	case fos.CountLines:
		line_count := fos.countMatchingLinesMap[filename]
		return fmt.Sprintf("%d", line_count)
	}
	return ""
}
