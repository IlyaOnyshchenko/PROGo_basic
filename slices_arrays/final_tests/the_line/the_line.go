package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var m int
	fmt.Scan(&m)
	n1 := 0
	for i := 0; i < n; i++ {
		if m < a[i] {
			continue
		}
		if a[i] < m {
			n1 = i - 1
			break
		}
		if a[i] == m {
			n1 = i
			continue
		}
	}
	fmt.Println(n1 + 2)
}
