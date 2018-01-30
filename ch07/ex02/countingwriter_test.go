package countingwriter

import (
	"os"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	data := []string{
		"HelloWolrd",
		"KIRIN LEMON",
		"Steel series",
	}
	w, c := CountingWriter(os.Stdout)
	var total int64 = 0
	for _, d := range data {
		bytes := []byte(d)
		w.Write(bytes)
		total += int64(len(bytes))

		if *c != total {
			t.Errorf("count shoud be %d count:%d\ntotal%d\n", total, *c, total)
		}
	}

}
