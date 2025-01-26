package main

import "fmt"

// package level variables
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boling point = %gF or %gC\n", f, c)
}
