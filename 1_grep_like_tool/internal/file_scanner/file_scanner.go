package grep_file_scanner

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

type FileScanner struct {
	Finder    Finder
	Paths     []string // can be both file paths or dir paths
	Recursive bool
	printOut  *io.Writer
	printErr  *io.Writer
}

func NewFileScanner(finder Finder, paths []string, recursive bool,
	print_out io.Writer, print_err io.Writer) FileScanner {

	var fs FileScanner
	fs.Finder = finder
	fs.Paths = paths
	fs.Recursive = recursive
	fs.printOut = &print_out
	fs.printErr = &print_err
	return fs
}

func (fileScanner FileScanner) GoThroughFiles() (string, error) {

	var workersCh = make(chan bool, workersNb)
	var resultCh = make(chan string)
	var workersWg = sync.WaitGroup{}
	var filesOut chan chan string = make(chan chan string, parallelFilesCollectNb)
	// fileOut can't be a list of chan string because we fill the list step by step and reassign it

	go fileScanner.collectOutput(filesOut, resultCh)

	// order of the results is important : for each file AND for each line
	workersWg.Add(1)

	if !fileScanner.Recursive {
		for _, filename := range fileScanner.Paths {
			fileScanner.processFile(filename, workersCh, filesOut, &workersWg)
		}

	} else {
		for _, filename := range fileScanner.Paths {

			visitor := func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() {
					fileScanner.processFile(path, workersCh, filesOut, &workersWg)
				}
				return nil
			}

			err := filepath.WalkDir(filename, visitor)

			if err != nil {
				err = fmt.Errorf("error while walking dir : %w", err)
				return "", err
			}
		}
	}
	workersWg.Done()
	workersWg.Wait()
	close(filesOut) // triggers return of collector (send to result)
	return <-resultCh, nil
}

func (fs FileScanner) collectOutput(filesOut chan chan string, result chan<- string) {
	var output string
	for fileCh := range filesOut {
		// cannot put this in a goroutine because we need to keep order here
		for line := range fileCh {
			output += line
		}
	}
	result <- output
}

func (fs FileScanner) processFile(filename string, workersCh chan bool, filesOut chan chan string, workersWg *sync.WaitGroup) {
	currentFileOut := make(chan string, parallelLinesCollectNb)
	filesOut <- currentFileOut
	workersCh <- true
	workersWg.Add(1)
	go fs.processFileConc(filename, currentFileOut, workersCh, workersWg)
}

func (fs FileScanner) processFileConc(filename string, currentFileOut chan<- string, workersCh <-chan bool, workersWg *sync.WaitGroup) {
	defer func() {
		close(currentFileOut)
		<-workersCh // free a worker spot
		workersWg.Done()
	}()

	file, err := os.Open(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open file %s because of %s", filename, err)
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			if utf8.ValidString(line) {
				currentFileOut <- fs.Finder.OutputOnLine(line, filename)
			} // else {
			// 	fmt.Fprintln(os.Stderr, "invalid line, not utf-8")
			// }

			// TODO : decide what to do with this
		}
		currentFileOut <- fs.Finder.OutputOnWholeFile(filename)
	}
}
