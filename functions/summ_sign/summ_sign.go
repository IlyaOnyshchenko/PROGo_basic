package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	sign := sign(a) + sign(b)
	fmt.Println(sign)
}

func sign(x int) int {
	var znac int
	if x > 0 {
		znac = 1
	}
	if x < 0 {
		znac = -1
	}
	if x == 0 {
		znac = 0
	}
	return znac
}
