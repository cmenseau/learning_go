package grep_output_control

type OutputControlRequest struct {
	CountLines        bool
	FilesWithoutMatch bool
	FilesWithMatch    bool
	OnlyMatching      bool
}
