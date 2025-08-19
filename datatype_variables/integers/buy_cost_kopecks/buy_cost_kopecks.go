package main

import (
	"fmt"
)

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	b = b * n
	a = a * n
	fmt.Println(a+b/100, b%100)
}
