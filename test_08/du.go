package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"flag"
)

func dirents(dir string) []os.FileInfo{
	entries, err := ioutil.ReadDir(dir)
	if err != nil{
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0{
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	for size := range fileSizes{
		nfiles++
		nbytes += size
	}
	fmt.Printf("%d files, %d\n", nfiles, nbytes)
}
