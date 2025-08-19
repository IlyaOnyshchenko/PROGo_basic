package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a)
	b = a % 10
	a = a / 10
	c = a % 10
	a = a / 10
	d = a % 10
	if b != c && c != d && b != d {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
