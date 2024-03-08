package grep_output_control

type FileOutputRequest struct {
	CountLines        bool
	FilesWithoutMatch bool
	FilesWithMatch    bool
}

type FileOutputSelect struct {
	FileOutputReq FileOutputRequest
	// TODO some data to save here to print recap when normal output is suppressed
}

func (fos FileOutputSelect) SuppressNormalOutput() bool {
	return fos.FileOutputReq.CountLines ||
		fos.FileOutputReq.FilesWithoutMatch ||
		fos.FileOutputReq.FilesWithMatch
}

func (fos FileOutputSelect) ProcessOutputLine(line string, filename string) {
	// TODO
	// store count or what's required to be able to print final output by file
}

func (fos FileOutputSelect) GetFinalOutputControl(filename string) string {
	// TODO
	return ""
}
