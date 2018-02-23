package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, wg *sync.WaitGroup, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	wg.Done()
}
func handleConn(c net.Conn) {
	conn := c.(*net.TCPConn)
	defer conn.CloseRead()
	input := bufio.NewScanner(conn)
	wg := &sync.WaitGroup{}
	ack := make(chan struct{})
	go timeoutTimer(ack, conn)
	for input.Scan() {
		wg.Add(1)
		ack <- struct{}{}
		go echo(conn, wg, input.Text(), 1*time.Second)
	}
	wg.Wait()
	conn.CloseWrite()
}

func timeoutTimer(ack chan struct{}, conn net.Conn) {
	for {
		select {
		case _, ok := <-ack:
			if !ok {
				return
			}

		case <-time.After(10 * time.Second):
			conn.Close()
			return
		}
	}
}
