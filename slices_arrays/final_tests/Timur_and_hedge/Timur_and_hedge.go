package main

import (
	"fmt"
)

func main() {
	var n, h int
	fmt.Scan(&n, &h)
	s := 0
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		if a[i] > h {
			s = s + 2
		} else {
			if a[i] <= h {
				s = s + 1
			}
		}
	}
	fmt.Println(s)
}
