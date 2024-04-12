package grep_output_control

import "fmt"

type FileOutputRequest struct {
	CountLines            bool
	FilesWithoutMatch     bool
	FilesWithMatch        bool
	filesWithoutMatchMap  map[string]bool
	filesWithMatchMap     map[string]bool
	countMatchingLinesMap map[string]int
}

func (fos FileOutputRequest) SuppressNormalOutput() bool {
	return fos.CountLines ||
		fos.FilesWithoutMatch ||
		fos.FilesWithMatch
}

func (fos *FileOutputRequest) ProcessOutputLine(line string, filename string) {
	// store count or what's required to be able to print final output by file

	if fos.FilesWithoutMatch {
		if fos.filesWithoutMatchMap == nil {
			fos.filesWithoutMatchMap = make(map[string]bool)
		}
		if line == "" {
			fos.filesWithoutMatchMap[filename] = true
		}
	}
	if fos.FilesWithMatch {
		if fos.filesWithMatchMap == nil {
			fos.filesWithMatchMap = make(map[string]bool)
		}
		if line != "" {
			fos.filesWithMatchMap[filename] = true
		}
	}
	if fos.CountLines {
		if fos.countMatchingLinesMap == nil {
			fos.countMatchingLinesMap = make(map[string]int)
		}
		if line != "" {
			fos.countMatchingLinesMap[filename] += 1
		}
	}
}

func (fos *FileOutputRequest) GetFinalOutputControl(filename string) string {
	if fos.FilesWithoutMatch {
		if fos.filesWithoutMatchMap[filename] {
			return filename
		}
	} else if fos.FilesWithMatch {
		if fos.filesWithMatchMap[filename] {
			return filename
		}
	} else if fos.CountLines {
		line_count := fos.countMatchingLinesMap[filename]
		return fmt.Sprintf("%d", line_count)
	}
	return ""
}
