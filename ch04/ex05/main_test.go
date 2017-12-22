package main

import (
	"fmt"
	"testing"
)

func main() {
	s := []string{"golang", "golang", "ruby", "hello", "hello", "ruby"}
	fmt.Println(deleteaAjacentDuplicate(s))
}
func TestdeleteaAjacentDuplicate(t *testing.T)
func deleteaAjacentDuplicate(s []string) []string {
	c := 0
	for i := 0; i < len(s)-1; i++ {
		if s[c] != s[i+1] {
			s[c] = s[i+1]
			c++
		}
	}
	return s
}
