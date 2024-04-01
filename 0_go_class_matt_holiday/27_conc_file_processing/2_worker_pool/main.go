package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

type pair struct {
	path string
	hash string
}

type fileList []string
type hashTable map[string]fileList

func hashFile(path string) pair {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hasher := md5.New()

	_, err = io.Copy(hasher, file)

	if err != nil {
		log.Fatal(err)
	}

	return pair{path: path, hash: string(hasher.Sum(nil))}
}

func searchTree(dir string, outputPaths chan<- string) (hashTable, error) {
	var res hashTable = make(hashTable, 0)

	visitor := func(path string, info fs.FileInfo, err error) error {
		//fmt.Println("visiting ", path)
		if err != nil && err != os.ErrNotExist {
			return err
		}

		if !info.IsDir() && info.Size() > 0 {
			outputPaths <- path
		}
		return nil
	}

	err := filepath.Walk(dir, visitor)

	return res, err
}

func workers(paths <-chan string, outPairs chan<- pair, done chan<- bool) {
	for path := range paths {
		outPairs <- hashFile(path)
	}
	done <- true
}

func collectHashes(pairs <-chan pair, result chan<- hashTable) {
	res := make(hashTable)
	for p := range pairs {
		res[p.hash] = append(res[p.hash], p.path)
	}
	result <- res
}

func run(dir string) (hashTable, error) {
	workersNb := 8
	paths := make(chan string)
	pairs := make(chan pair)
	done := make(chan bool)
	results := make(chan hashTable)

	// start 8 goroutines
	// will be blocked until paths are fed into paths channel
	for i := 0; i < workersNb; i++ {
		go workers(paths, pairs, done)
	}

	// start collector goroutine
	// will start collecting once it receives pairs from pairs channel from workers
	go collectHashes(pairs, results)

	// start feeding paths to the workers
	// not concurrent
	searchTree(dir, paths)

	// close paths to let workers now that this is the end of their tasks
	// workers may not have finished at that time
	close(paths)

	// wait for all workers to be done
	// terminate all goroutines
	for i := 0; i < workersNb; i++ {
		<-done
	}

	// all workers goroutines are finished : we can close pairs
	// it will trigger collector to send final hashtable result
	close(pairs)

	res := <-results

	return res, nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Missing argument path")
	}

	rootPath := os.Args[1]

	start := time.Now()

	res, err := run(rootPath)

	fmt.Println(time.Since(start))

	if err != nil {
		log.Fatal(err)
	}

	for _, list := range res {
		if len(list) > 1 {
			fmt.Println("Same file content detected for ", list)
		}
	}

}
