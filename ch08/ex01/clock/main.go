package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8000", "listen port")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //例: 接続が切れた
			continue
		}
		go handleConn(conn) //一度にひとつの接続を処理する
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		//		_, err := io.WriteString(c, "\033[A\033[2K\r"+time.Now().Format("15:04:05\n"))
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
