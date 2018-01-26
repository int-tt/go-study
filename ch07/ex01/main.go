package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type LineCounter int

func main() {
	var str string
	var line LineCounter
	if len(os.Args[1]) != 0 {
		str = os.Args[1]
	}
	i, err := line.Write([]byte(str))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\tCount:%d\n", str, i)
}
func (c *LineCounter) Write(p []byte) (int, error) {
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
