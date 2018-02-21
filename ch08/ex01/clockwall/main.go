package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatal("You need to specify at least one server")
	}
	clocks := make(map[string]string, len(os.Args[1:]))
	for _, arg := range os.Args[1:] {
		clock := strings.Split(arg, "=")
		clocks[clock[0]] = clock[1]
	}
	for location, host := range clocks {
		conn, err := net.Dial("tcp", host)
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(location, conn)
	}

	for {
	}
}
func handleConnection(location string, conn net.Conn) {
	defer conn.Close()
	for {
		buf := bufio.NewReader(conn)
		line, _, err := buf.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(os.Stdout, fmt.Sprintf("%s: %s\n", location, string(line)))
	}
}

// func mustCopy(dst io.Writer, src io.Reader) {
// 	if _, err := io.Copy(dst, src); err != nil {
// 		log.Fatal(err)
// 	}
// }
