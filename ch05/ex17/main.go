package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	u := os.Args[1]
	resp, err := http.Get(u)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ElmentsByTagName(doc, "img"))
	fmt.Println(ElmentsByTagName(doc, "a", "h1"))
}

func ElmentsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node

	if doc.Type == html.ElementNode {
		for _, tagName := range name {
			if tagName == doc.Data {
				nodes = append(nodes, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		for _, node := range ElmentsByTagName(c, name...) {
			nodes = append(nodes, node)
		}
	}
	return nodes
}
