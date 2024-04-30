package file_scanner

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unicode/utf8"
)

const workersNb int = 8
const parallelFilesCollectNb = 10
const parallelLinesCollectNb = 20

type Finder interface {
	OutputOnLine(line string, filename string) string
	OutputOnWholeFile(filename string) string
}

type FileSelectionContainer interface {
	IsRecursive() bool
	GetPaths() []string // can be both file paths or dir paths
}

type FileScanner struct {
	Finder                 Finder
	FileSelectionContainer FileSelectionContainer
	printOut               *io.Writer
	printErr               *io.Writer
}

func NewFileScanner(finder Finder, fsc FileSelectionContainer,
	print_out io.Writer, print_err io.Writer) FileScanner {

	var fs FileScanner
	fs.Finder = finder
	fs.FileSelectionContainer = fsc
	fs.printOut = &print_out
	fs.printErr = &print_err
	return fs
}

type typedMsg struct {
	content string
	isErr   bool
}

func (fileScanner FileScanner) GoThroughFiles() {

	var workersCh = make(chan bool, workersNb)
	var resultCh = make(chan bool)
	var workersWg = sync.WaitGroup{}
	var filesOut chan chan typedMsg = make(chan chan typedMsg, parallelFilesCollectNb)
	// fileOut can't be a list of chan string because we fill the list step by step and reassign it

	go fileScanner.collectOutput(filesOut, resultCh)

	// order of the results is important : for each file AND for each line
	workersWg.Add(1)

	if !fileScanner.FileSelectionContainer.IsRecursive() {
		for _, filename := range fileScanner.FileSelectionContainer.GetPaths() {

			fileScanner.processFile(filename, workersCh, filesOut, &workersWg)
		}

	} else {
		for _, filename := range fileScanner.FileSelectionContainer.GetPaths() {

			visitor := func(path string, d fs.DirEntry, err error) error {

				if err != nil {
					fileScanner.processError(filesOut, err)
					return nil
				}
				if !d.IsDir() {
					fileScanner.processFile(path, workersCh, filesOut, &workersWg)
				}
				return nil
			}

			err := filepath.WalkDir(filename, visitor)
			if err != nil {
				fileScanner.processError(filesOut, err)
			}
		}
	}
	workersWg.Done()
	workersWg.Wait()
	close(filesOut) // triggers return of collector (send to result)
	<-resultCh
}

func (fs FileScanner) collectOutput(filesOut chan chan typedMsg, result chan<- bool) {
	var outBuffer strings.Builder
	for fileCh := range filesOut {
		// cannot put this in a goroutine because we need to keep order here
		for msg := range fileCh {
			if msg.isErr {
				// flush out, to print error
				fmt.Fprint(*fs.printOut, outBuffer.String())
				outBuffer.Reset()
				fmt.Fprintln(*fs.printErr, msg.content)
			} else {
				outBuffer.WriteString(msg.content)
			}
		}
	}
	if outBuffer.Len() != 0 {
		fmt.Fprint(*fs.printOut, outBuffer.String())
	}
	result <- true
}

func (fs FileScanner) processFile(filename string, workersCh chan bool, filesOut chan chan typedMsg, workersWg *sync.WaitGroup) {
	currentFileOut := make(chan typedMsg, parallelLinesCollectNb)
	filesOut <- currentFileOut
	workersCh <- true
	workersWg.Add(1)
	go fs.processFileConc(filename, currentFileOut, workersCh, workersWg)
}

func (fs FileScanner) processError(filesOut chan chan typedMsg, err error) {
	currentFileOut := make(chan typedMsg, parallelLinesCollectNb)
	filesOut <- currentFileOut
	currentFileOut <- typedMsg{
		content: err.Error(),
		isErr:   true,
	}
	close(currentFileOut)
}

func (fs FileScanner) processFileConc(filename string, currentFileOut chan typedMsg, workersCh <-chan bool, workersWg *sync.WaitGroup) {
	defer func() {
		close(currentFileOut)
		<-workersCh // free a worker spot
		workersWg.Done()
	}()

	file, err := os.Open(filename)

	if err != nil {
		currentFileOut <- typedMsg{
			content: err.Error(),
			isErr:   true,
		}
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if utf8.ValidString(line) {
			currentFileOut <- typedMsg{
				content: fs.Finder.OutputOnLine(line, filename),
				isErr:   false,
			}
		} // else {
		// 	fmt.Fprintln(os.Stderr, "invalid line, not utf-8")
		// }

		// TODO : add support of -a option
	}
	currentFileOut <- typedMsg{
		content: fs.Finder.OutputOnWholeFile(filename),
		isErr:   false,
	}
}
