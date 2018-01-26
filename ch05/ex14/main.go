package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

var initialURL *url.URL

func main() {
	breadthFirst(listFiles, os.Args[1:])
}

// breadThFirstはworklist内の個々の項目に対してfを呼び出します。
// fからかえされたすべての項目はworklistへ追加されます。
// fは、それぞれの項目に対して高々一度しか呼び出されません。
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func listFiles(path string) []string {
	fmt.Println(path)
	dirInfos := extractFileInfos(path)

	var files []string
	for _, dirInfo := range dirInfos {
		name := dirInfo.Name()
		if name[0] == '.' {
			continue
		}
		files = append(files, path+"/"+dirInfo.Name())
	}
	return files
}
func extractFileInfos(path string) []os.FileInfo {
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		log.Println(err)
		return nil
	}

	if !fileInfo.IsDir() {
		return nil
	}

	dirInfos, err := f.Readdir(0)
	if err != nil {
		log.Println(err)
	}
	return dirInfos

}
