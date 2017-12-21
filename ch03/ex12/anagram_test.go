package main

import (
	"testing"
)

func TestAnagram(t *testing.T) {
	data := map[string]string{
		"golang":  "langgo",
		"integer": "egerint",
		"float":   "taolf",
		"ruby":    "bury",
	}
	for input, expeted := range data {
		if !isAnagram(input, expeted) {
			t.Errorf("%s and %s are not aganrams", input, expeted)
		}
	}
}
