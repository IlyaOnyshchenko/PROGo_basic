package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if (raz(a) == true) && (raz(b) == true) {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}

func raz(x int) bool {
	var s, s1 int
	var flag bool
	for i := 6; i > 3; i-- {
		s = s + x%10
		x = x / 10
	}
	for i := 3; i > 0; i-- {
		s1 = s1 + x%10
		x = x / 10
	}
	if s == s1 {
		flag = true
	}
	return flag
}
