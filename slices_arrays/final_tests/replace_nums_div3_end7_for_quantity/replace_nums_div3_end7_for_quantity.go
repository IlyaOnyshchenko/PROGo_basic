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
	for i := 0; i < n; i++ {
		if (a[i]%3 == 0) && (a[i]%10 == 7) {
			s++
		}
	}
	for i := 0; i < n; i++ {
		if (a[i]%3 == 0) && (a[i]%10 == 7) {
			a[i] = s
		}
	}
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
}
