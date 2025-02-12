package main

import "fmt"

var n int

func init() {
	fmt.Println("Func init")
	n = 10
}

func main() {
	fmt.Println("main")
	fmt.Println(n)
}
