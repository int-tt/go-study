package LimitReader

import (
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	data := []string{
		"",
		"Hello",
		"Hello Word",
		"HelloWorld",
	}
	for _, d := range data {
		str := strings.NewReader(d)
		lr := LimitReader(str, 3)
		n, err := lr.Read([]byte("test"))
		t.Log(n, err)
	}
}
