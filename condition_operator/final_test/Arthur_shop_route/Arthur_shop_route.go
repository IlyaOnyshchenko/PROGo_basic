package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var s1, s2, s3, s4 int
	fmt.Scan(&a, &b, &c)
	s1 = a + b + c
	s2 = 2*a + 2*b
	s3 = 2*a + 2*c
	s4 = 2*b + 2*c
	fmt.Println(min(s1, s2, s3, s4))
}
