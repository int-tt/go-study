package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello World"
	fmt.Println(string(reverseUTF8([]byte(s))))
}
func reverseUTF8(bytes []byte) []byte {
	var size int
	for i := 0; i < len(bytes)/2; i += size {
		_, size = utf8.DecodeRune(bytes)
		for j, k := 0, len(bytes)-1; j < k; j, k = j+1, k-1 {
			bytes[j], bytes[k] = bytes[k], bytes[j]
		}
	}
	return bytes
}
