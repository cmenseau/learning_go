package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
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

func searchTree(dir string, wg *sync.WaitGroup, pairs chan<- pair, limits chan bool) (hashTable, error) {
	var res hashTable = make(hashTable, 0)
	defer wg.Done()

	visitor := func(path string, info fs.FileInfo, err error) error {
		//fmt.Println("visiting ", path)
		if err != nil && err != os.ErrNotExist {
			return err
		}

		if info.IsDir() && dir != path {
			wg.Add(1)
			go searchTree(path, wg, pairs, limits)
			return filepath.SkipDir
		} else if !info.IsDir() && info.Size() > 0 {
			wg.Add(1)
			go worker(path, pairs, wg, limits)
		}
		return nil
	}

	err := filepath.Walk(dir, visitor)

	return res, err
}

func worker(path string, outPairs chan<- pair, wg *sync.WaitGroup, limits chan bool) {
	defer wg.Done()
	limits <- true
	outPairs <- hashFile(path)
	<-limits
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
	pairs := make(chan pair)
	results := make(chan hashTable)
	limits := make(chan bool, workersNb)
	var wgDir sync.WaitGroup

	// start collector goroutine
	// will start collecting once it receives pairs from pairs channel from workers
	go collectHashes(pairs, results)

	// start feeding paths to the workers
	// not concurrent
	wgDir.Add(1)
	searchTree(dir, &wgDir, pairs, limits)
	//wgDir.Done()

	wgDir.Wait()

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
