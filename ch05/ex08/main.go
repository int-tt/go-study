package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	//TODO: flag対応したい
	if len(os.Args) != 3 {
		fmt.Println("usage: findElementName uri id")
	}
	findElement(os.Args[1], os.Args[2])
}

func findElement(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	node := ElementByID(doc, id)
	if node == nil {
		fmt.Printf("no element with \"%s\" attribute Found\n ", id)
	} else {
		printNode(node)
	}
	return nil
}
func printNode(n *html.Node) {
	fmt.Printf("<%s", n.Data)
	for _, v := range n.Attr {
		fmt.Printf(" %s='%s'", v.Key, v.Val)
	}
	fmt.Println(">")
}

// forEachはnから始まるツリー無いの個々のノードxに対して
// 関数pre(x)とpost(x)を呼び出します。その2つの関数はオプションです
// preは子ノードを訪れる前に呼び出され(全順preorder)
// postは子ノードを訪れたあとに呼び誰覚ます(降順:postorder)
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}
	if post != nil {
		if !post(n) {
			return false
		}
	}
	return true
}
func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node

	forEachNode(doc, func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, v := range n.Attr {
			if v.Val == id && v.Key == "id" {
				node = n
				return false
			}
		}
		return true
	}, nil)

	return node
}

var depth int

func startElement(n *html.Node) {
	depth++
	if n.FirstChild == nil {
		return
	}
	attribute := getAttribute(n.Attr)
	if attribute == "" {
		fmt.Printf("\n%*s<%s>", depth, "", n.Data)
	} else {
		fmt.Printf("\n%*s<%s %s>", depth, "", n.Data, attribute)
	}
}

func endElement(n *html.Node) {
	if n.FirstChild == nil {
		attribute := getAttribute(n.Attr)
		if attribute == "" {
			switch n.Data {
			case "br":
				fmt.Printf("<%s>\n", n.Data)
			default:
				fmt.Printf("\n%*s<%s %s>", depth, "", n.Data, attribute)
			}
		} else {
			fmt.Printf("\n%*s<%s %s />", depth, "", n.Data, attribute)
		}
	} else {
		switch n.Data {
		case "a", "code", "title", "tt", "h1":
			fmt.Printf("</%s>", n.Data)
		default:
			fmt.Printf("\n%*s</%s>", depth, "", n.Data)
		}
	}
	depth--
}
func startNode(n *html.Node) {
	switch n.Type {
	case html.ErrorNode:
	case html.TextNode:
		startTextNode(n)
		return
	case html.DocumentNode:
	case html.ElementNode:
		startElement(n)
		return
	case html.CommentNode:
	case html.DoctypeNode:
	}
}
func endNode(n *html.Node) {
	switch n.Type {
	case html.ErrorNode:
	case html.TextNode:
		return
	case html.DocumentNode:
	case html.ElementNode:
		endElement(n)
		return
	case html.CommentNode:
	case html.DoctypeNode:
	}
}
func startTextNode(n *html.Node) {
	fmt.Printf("%s", n.Data)
}

func getAttribute(attribute []html.Attribute) string {
	var buf bytes.Buffer

	for k, v := range attribute {
		if k != 0 {
			buf.WriteString(" ")
		}

		if v.Namespace == "" {
			buf.WriteString(v.Key)
			buf.WriteString(`="`)
			buf.WriteString(v.Val)
			buf.WriteString(`"`)
		} else {
			buf.WriteString(v.Namespace)
			buf.WriteString(":")
			buf.WriteString(v.Key)
			buf.WriteString(`="`)
			buf.WriteString(v.Val)
			buf.WriteString(`"`)
		}
	}
	return buf.String()
}

func printDocType(n *html.Node) {
	if n.Type != html.DoctypeNode {
		//panic("Illegal Argument")
		return
	}
	var buf bytes.Buffer

	buf.WriteString("<!DOCTYPE ")
	buf.WriteString(n.Namespace)

	for k, v := range n.Attr {
		if k != 0 {
			buf.WriteString(" ")
		}

		if v.Key == "public" {
			buf.WriteString("PUBLIC ")
			buf.WriteString(`"`)
		}
	}

}
