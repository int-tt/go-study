package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, text := range os.Args[1:] {
		doc, err := html.Parse(NewReader(text))
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1:%c\n", err)
			os.Exit(1)
		}
		for _, links := range visit(nil, doc) {
			fmt.Println(links)
		}
	}
}

type reader struct {
	s string
	i int
}

func (r *reader) Read(p []byte) (int, error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n := copy(p, r.s[r.i:])
	r.i += n
	return n, nil
}

func NewReader(s string) io.Reader {
	return &reader{s, 0}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
