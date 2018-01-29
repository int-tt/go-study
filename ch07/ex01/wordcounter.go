package wordcounter

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	nBytes := len(p)

	scanner := bufio.NewScanner(bytes.NewReader(p))

	for scanner.Scan() {
		_ = scanner.Text()
		*c += 1
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("scanner error %v", err))
	}

	return nBytes, nil
}
