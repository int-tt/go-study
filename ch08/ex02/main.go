package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//21は制御用,20は転送用
func main() {
	listner, err := net.Listen("tcp", "localhost:8021")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer closeConn(conn)
	log.Println("Connected.")
	exDir, err := os.Executable()
	if err != nil {

	}
	exPath := filepath.Dir(exDir)
	buf := bufio.NewScanner(conn)
	for buf.Scan() {
		line := buf.Text()
		log.Println(line)
		args := strings.Split(string(line), " ")
		switch args[0] {
		case "CWD":
		case "QUIT":
		case "PWD":
			io.WriteString(conn, exPath)
		case "LIST":
			out, err := exec.Command("ls", "-lah", exPath).Output()
			if err != nil {
				log.Println(err)
				continue
			}
			io.WriteString(conn, string(out))
		case "RETR":
		default:
		}
	}
}

func closeConn(conn net.Conn) {
	defer conn.Close()
	_, err := io.WriteString(conn, "Connection close.")
	if err != nil {
		log.Println(err)
	}
	log.Println("Connection close.")
}
