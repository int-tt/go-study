package main

import (
	"bytes"
	"fmt"
)

func main() {
	vals := []string{"one", "two", "three"}
	fmt.Println(join(",", vals...))
}

func join(sep string, vals ...string) string {
	if len(vals) == 0 {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString(vals[0])
	for _, v := range vals[1:] {
		buf.WriteString(sep)
		buf.WriteString(v)
	}
	return buf.String()
}
