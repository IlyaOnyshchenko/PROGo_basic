package main

import (
	"fmt"
)

func main() {
	var n, s int
	fmt.Scan(&n)
	s = n%10 + s
	n = n / 10
	s = n%10 + s
	n = n / 10
	s = n%10 + s
	fmt.Println(s)
}
