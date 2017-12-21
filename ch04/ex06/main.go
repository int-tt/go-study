package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "Hello   Wor  ld"
	fmt.Println(string(combiendSpace([]byte(s))))
}
func combiendSpace(bytes []byte) []byte {
	whiteSpaceBuf := make([]byte, 4)
	whiteSpaceSize := utf8.EncodeRune(whiteSpaceBuf, ' ')
	whiteSpaceBuf = whiteSpaceBuf[:whiteSpaceSize]

	c := 0
	var size int
	var r rune
	isSpace := false
	for i := 0; i < len(bytes)-1; i += size {
		r, size = utf8.DecodeRune(bytes[i:])

		if unicode.IsSpace(r) {
			if !isSpace {
				copy(bytes[c:], whiteSpaceBuf)
				c += whiteSpaceSize
				isSpace = true
			}
			continue
		}

		//UTF8 1文字分だけコピー
		copy(bytes[c:], bytes[i:i+size])
		c += size
		isSpace = false
	}

	return bytes[:c]
}
