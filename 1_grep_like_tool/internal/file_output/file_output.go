package file_output

import "fmt"

type FileOutputRequest struct {
	CountLines            bool
	FilesWithoutMatch     bool
	FilesWithMatch        bool
	filesWithoutMatchMap  map[string]bool // set of files without match
	filesWithMatchMap     map[string]bool // set of files with matches
	countMatchingLinesMap map[string]int
}

type FileOutputSelector struct {
	Fo *FileOutputRequest
}

func (fos FileOutputSelector) suppressNormalOutput() bool {
	return fos.Fo.CountLines ||
		fos.Fo.FilesWithoutMatch ||
		fos.Fo.FilesWithMatch
}

func (fos *FileOutputSelector) ProcessLine(line string, filename string) string {
	// store count or what's required to be able to print final output by file

	// when we don't process the results at the file level
	if !fos.suppressNormalOutput() {
		return line
	}

	switch {
	case fos.Fo.FilesWithoutMatch:
		if fos.Fo.filesWithoutMatchMap == nil {
			fos.Fo.filesWithoutMatchMap = make(map[string]bool)
		}
		if line == "" {
			fos.Fo.filesWithoutMatchMap[filename] = true
		}
	case fos.Fo.FilesWithMatch:
		if fos.Fo.filesWithMatchMap == nil {
			fos.Fo.filesWithMatchMap = make(map[string]bool)
		}
		if line != "" {
			fos.Fo.filesWithMatchMap[filename] = true
		}
	case fos.Fo.CountLines:
		if fos.Fo.countMatchingLinesMap == nil {
			fos.Fo.countMatchingLinesMap = make(map[string]int)
		}
		if line != "" {
			fos.Fo.countMatchingLinesMap[filename] += 1
		}
	}
	return ""
}

func (fos *FileOutputSelector) GetFileLevelResult(filename string) string {
	switch {
	case fos.Fo.FilesWithoutMatch:
		if fos.Fo.filesWithoutMatchMap[filename] {
			return filename
		}
	case fos.Fo.FilesWithMatch:
		if fos.Fo.filesWithMatchMap[filename] {
			return filename
		}
	case fos.Fo.CountLines:
		line_count := fos.Fo.countMatchingLinesMap[filename]
		return fmt.Sprintf("%d", line_count)
	}
	return ""
}
