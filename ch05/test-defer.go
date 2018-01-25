package main

import (
	"fmt"
	"os"
	"time"
)

var TempDir = "tmp"

func main() {

	_ = os.MkdirAll(TempDir, os.ModePerm)
	time.Sleep(3 * time.Second)
	defer removeTempDir()
	time.Sleep(3 * time.Second)
}
func removeTempDir() {
	_ = os.Remove(TempDir)
	fmt.Println("deleted!")
}
