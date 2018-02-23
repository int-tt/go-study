package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gopl/ch05/links"
)

type node struct {
	depth int
	links []string
}

var tokens = make(chan struct{}, 20)
var limitDepth = flag.Int("depth", 2, "depth (default:2)")

func main() {

	flag.Parse()

	worklist := make(chan *node)
	var n int

	n++
	go func() {
		worklist <- &node{0, flag.Args()}
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		nodes := <-worklist
		for _, link := range nodes.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(depth int, link string) {
					worklist <- crawl(depth, link)
				}(nodes.depth, link)
			}
		}
	}
}

func crawl(depth int, u string) *node {
	if depth >= *limitDepth {
		return &node{depth + 1, nil}
	}
	fmt.Println(u)
	tokens <- struct{}{}
	list, err := links.Extract(u)
	<-tokens

	if err != nil {
		log.Println(err)
	}
	return &node{depth + 1, list}
}
