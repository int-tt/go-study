package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, u := range os.Args[1:] {
		local, n, err := fetch(u)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", u, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes.) \n", u, local, n)
	}
}

func fetch(u string) (filename string, n int64, err error) {
	resp, err := http.Get(u)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		// derferが異常終了時に呼ばれた場合,errを上書きしないようにする
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
