package main

import (
	"flag"
	"fmt"

	"github.com/int-tt/go-study/ch07/tempconv/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}