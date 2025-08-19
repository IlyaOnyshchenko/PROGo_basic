package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e, f, g int
	fmt.Scan(&a)
	b = a % 10
	a /= 10
	c = a % 10
	a /= 10
	d = a % 10
	a /= 10
	e = a % 10
	a /= 10
	f = a % 10
	a /= 10
	g = a % 10
	if d+b+c == e+f+g {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
