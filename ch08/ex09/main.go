package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	log.Println(roots)
	if len(roots) == 0 {
		roots = []string{"."}
	}
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go run(root, &n, os.Stdout)
	}
	n.Wait()
}
func run*dir string,n *sync.WaitGroup,w io.Writer){
	defer n.Done()
	fileSize := make(chan int64)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go walkDir(dir,wg,fileSizes)
}
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
