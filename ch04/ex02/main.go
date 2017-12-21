package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	shaType := flag.String("a", "256", `"256" "384" "512"`)
	flag.Parse()

	if *shaType != "256" && *shaType != "384" && *shaType != "512" {
		log.Fatalf("Invalid Algorithm: sha%s\n", *shaType)
	}
	reader := bufio.NewReader(os.Stdin)
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal("error:", err)
	}
	switch *shaType {
	case "256":
		sum := sha256.Sum256(bytes)
		fmt.Printf("%x\n", sum)
	case "384":
		sum := sha512.Sum384(bytes)
		fmt.Printf("%x\n", sum)
	case "512":
		sum := sha512.Sum512(bytes)
		fmt.Printf("%x\n", sum)
	}

}
