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

	var outputCh = make(chan string, 10)
	var workersCh = make(chan bool, 5)
	var resultCh = make(chan string)
	var workersWg = sync.WaitGroup{}

	go fileScanner.collectOutput(outputCh, resultCh)

	// TODO :
	// order of the results is important : for each file AND for each line

	if !fileScanner.Recursive {
		for _, filename := range fileScanner.Paths {
			workersCh <- true
			workersWg.Add(1)
			go fileScanner.processFile(filename, outputCh, workersCh, &workersWg)
		}
	} else {
		for _, filename := range fileScanner.Paths {

			visitor := func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() {
					workersCh <- true
					workersWg.Add(1)
					go fileScanner.processFile(path, outputCh, workersCh, &workersWg)
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

	// handle output now

	workersWg.Wait()
	close(outputCh) // triggers return of collector (send to result)

	return <-resultCh, nil
}

func (fs FileScanner) collectOutput(matchingLines <-chan string, result chan<- string) {
	var output string
	for line := range matchingLines {
		output += line
	}
	result <- output
}

func (fs FileScanner) processFile(filename string, output chan<- string, workersCh <-chan bool, workersWg *sync.WaitGroup) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open file %s because of %s", filename, err)
	} else {

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			if utf8.ValidString(line) {
				output <- fs.Finder.OutputOnLine(line, filename)
			} // else {
			// 	fmt.Fprintln(os.Stderr, "invalid line, not utf-8")
			// }

			// TODO : decide what to do with this
		}
	}
	file.Close()
	output <- fs.Finder.OutputOnWholeFile(filename)
	<-workersCh
	workersWg.Done()
}
