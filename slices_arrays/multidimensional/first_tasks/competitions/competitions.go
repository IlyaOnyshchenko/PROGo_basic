package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	s := 0
	sMax := 0
	n1 := 0
	for i := 0; i < n; i++ {
		s = 0
		for j := 0; j < m; j++ {
			s = s + a[i][j]
		}
		if sMax < s {
			sMax = s
			n1 = i
		}
	}
	fmt.Println(sMax)
	fmt.Println(n1)
}
