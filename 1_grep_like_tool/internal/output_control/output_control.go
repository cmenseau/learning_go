package grep_output_control

import "fmt"

type FileOutputRequest struct {
	CountLines        bool
	FilesWithoutMatch bool
	FilesWithMatch    bool
}

type FileOutputSelect struct {
	FileOutputReq         FileOutputRequest
	filesWithoutMatchMap  map[string]bool
	filesWithMatchMap     map[string]bool
	countMatchingLinesMap map[string]int
}

func (fos *FileOutputSelect) SuppressNormalOutput() bool {
	return fos.FileOutputReq.CountLines ||
		fos.FileOutputReq.FilesWithoutMatch ||
		fos.FileOutputReq.FilesWithMatch
}

func (fos *FileOutputSelect) ProcessOutputLine(line string, filename string) {
	// store count or what's required to be able to print final output by file

	// order of the IF matters,
	// grep supports 1 option at a time in this order
	// TODO : enforce this logic in parser rather than here ?
	if fos.FileOutputReq.FilesWithoutMatch {
		if fos.filesWithoutMatchMap == nil {
			fos.filesWithoutMatchMap = make(map[string]bool)
		}
		if line == "" {
			fos.filesWithoutMatchMap[filename] = true
		}
	} else if fos.FileOutputReq.FilesWithMatch {
		if fos.filesWithMatchMap == nil {
			fos.filesWithMatchMap = make(map[string]bool)
		}
		if line != "" {
			fos.filesWithMatchMap[filename] = true
		}
	} else if fos.FileOutputReq.CountLines {
		if fos.countMatchingLinesMap == nil {
			fos.countMatchingLinesMap = make(map[string]int)
		}
		if line != "" {
			fos.countMatchingLinesMap[filename] += 1
		}
	}
}

func (fos *FileOutputSelect) GetFinalOutputControl(filename string) string {
	if fos.FileOutputReq.FilesWithoutMatch {
		if fos.filesWithoutMatchMap[filename] {
			return filename
		}
	} else if fos.FileOutputReq.FilesWithMatch {
		if fos.filesWithMatchMap[filename] {
			return filename
		}
	} else if fos.FileOutputReq.CountLines {
		return fmt.Sprintf("%d", fos.countMatchingLinesMap[filename])
	}
	return ""
}
