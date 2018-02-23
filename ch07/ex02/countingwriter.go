package countingwriter

import (
	"io"
)

type Counting struct {
	w io.Writer
	c int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cnt Counting
	cnt.w = w
	return &cnt, &(cnt.c)
}

func (c *Counting) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	if err != nil {
		return 0, err
	}
	c.c += int64(n)
	return
}
