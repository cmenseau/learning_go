package main

import (
	"fmt"
	grep_engine "main/internal/engine"
	grep_file_scanner "main/internal/file_scanner"
	grep_parser "main/internal/parser"
	"os"
)

func main() {
	req := grep_parser.ParseArgs(os.Args[1:])

	eng := grep_engine.Engine{Request: &req}

	fmt.Print(grep_file_scanner.GoThroughFiles(
		eng.Request.Filenames, eng.OutputOnLine, eng.OutputOnWholeFile))
}
