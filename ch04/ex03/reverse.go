package main

import "fmt"

func main() {
	str := "helloworld"
	rev := reverse([]byte(str))
	fmt.Println(string(rev))
}

func reverse(input []byte) []byte {
	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-1-i] = input[len(input)-1-i], input[i]
	}
	return input
}
