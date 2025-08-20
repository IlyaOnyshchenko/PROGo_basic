package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	sign := raz(a) * raz(b)
	fmt.Println(sign)
}

func raz(x int) int {
	var s int
	for x > 0 {
		x = x / 10
		s++
	}
	return s
}
