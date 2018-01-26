package main

import (
	"fmt"
	"os"
	"regexp"
)

var pattern = regexp.MustCompile(`(\$\w*)`)

func main() {
	fmt.Println(expand(os.Args[1], func(s string) string {
		return fmt.Sprintf("Hello! %s ", s)
	}))
}
func expand(s string, f func(string) string) string {
	r := pattern.ReplaceAllStringFunc(s, func(match string) string {
		return f(match[1:])
	})
	return r
}
