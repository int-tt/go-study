package main

import (
	"bytes"
	"io"
	"testing"
)

func TestParser(t *testing.T) {
	data := []string{
		"",
		"Hello",
		"Hello Word",
		"HelloWorld",
	}

	for _, d := range data {
		var b bytes.Buffer
		r := NewReader(d)
		n, err := io.Copy(&b, r)
		if err != nil {
			t.Error(err)
			continue
		}
		if except := int(n); except != len(d) {
			t.Errorf("%v,%v", n, len(d))
		}
	}
}
