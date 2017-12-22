package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456789"))
}

func comma(s string) string {
	var buf bytes.Buffer
	runes := []byte(s)
	cnt := len(runes) % 3
	if cnt == 0 {
		cnt = 3
	}
	for _, rune := range runes {

		if cnt == 0 {
			buf.WriteByte(',')
			cnt = 3
		}
		buf.WriteByte(rune)
		cnt--
	}
	return buf.String()
}
