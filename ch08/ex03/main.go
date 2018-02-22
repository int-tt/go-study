package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan struct{})
	tcpConn := conn.(*net.TCPConn)
	go func() {
		defer tcpConn.CloseRead()
		io.Copy(os.Stdout, tcpConn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(tcpConn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
