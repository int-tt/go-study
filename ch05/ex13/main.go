package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"gopl.io/ch5/links"
)

var initialURL *url.URL

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: crawl <url>")
		os.Exit(1)
	}
	var err error

	initialURL, err = url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(*initialURL)
	}

	breadthFirst(crawl, []string{os.Args[1]})
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
				if isSameDomain(item) {
					download(item)
					worklist = append(worklist, f(item)...)
				}
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func isSameDomain(item string) bool {
	u, err := url.Parse(item)
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}

	return strings.HasSuffix(u.Host, initialURL.Host)
}

func download(item string) {
	resp, err := http.Get(item)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	dir := path.Dir(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	if strings.HasSuffix(item, "/") {
		if strings.HasSuffix(dir, local) {
			local = "index.html"
		}
	}

	fmt.Printf("cached/%s%s %s\n", resp.Request.URL.Host, dir, local)
	fullDir := "cached/" + resp.Request.URL.Host + dir
	if err = os.MkdirAll(fullDir, os.ModePerm); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	file, err := os.Create(fullDir + "/" + local)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
