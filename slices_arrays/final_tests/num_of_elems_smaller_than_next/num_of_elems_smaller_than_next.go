package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := 0
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] {
			s++
		}
	}
	fmt.Println(s)
}
