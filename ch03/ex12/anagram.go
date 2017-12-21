package main

import (
	"bytes"
	"fmt"
	"sort"
)

type Byte []byte

func (b Byte) Len() int {
	return len(b)
}
func (b Byte) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b Byte) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]

}
func main() {
	s1 := "golang"
	s2 := "langgoo"
	fmt.Println(isAnagram(s1, s2))
}

func isAnagram(s1, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	sort.Sort(Byte(b1))
	sort.Sort(Byte(b2))
	return bytes.Compare(b1, b2) == 0
}
