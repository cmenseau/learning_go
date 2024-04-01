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

func searchTree(dir string) (hashTable, error) {
	var res hashTable = make(hashTable, 0)

	visitor := func(path string, info fs.FileInfo, err error) error {
		//fmt.Println("visiting ", path)
		if err != nil && err != os.ErrNotExist {
			return err
		}

		if !info.IsDir() && info.Size() > 0 {
			p := hashFile(path)
			res[p.hash] = append(res[p.hash], p.path)
		}
		return nil
	}

	err := filepath.Walk(dir, visitor)

	return res, err
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Missing argument path")
	}

	rootPath := os.Args[1]

	start := time.Now()

	res, err := searchTree(rootPath)

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
