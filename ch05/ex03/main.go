package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	count := make(map[string]int)
	countTag(count, doc)
	fmt.Printf("html tags\tcount\n")
	for k, v := range count {
		fmt.Printf("%-10s\t%d\n", k, v)
	}
}

func countTag(count map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && !strings.Contains("script", n.Data) && !strings.Contains("style", n.Data) {
		count[n.Data]++
	}
	countTag(count, n.FirstChild)
	countTag(count, n.NextSibling)

}
