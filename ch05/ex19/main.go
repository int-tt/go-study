package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
}
func add(i, j int) (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = p.(int)
		}
	}()
	panic(i + j)
}
