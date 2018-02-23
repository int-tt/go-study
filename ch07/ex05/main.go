package LimitReader

import "io"

type LimitedReader struct {
	r io.Reader
	l int64
	n int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n, 0}
}

func (l *LimitedReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if l.n >= l.l {
		return 0, io.EOF
	}
	nbytes := int(l.l - l.n)
	if int64(len(p)) > l.n {
		p = p[0:nbytes]
	}
	n, err := l.r.Read(p)
	if err != nil {
		return n, err
	}
	l.n += int64(nbytes)
	return n, nil
}
