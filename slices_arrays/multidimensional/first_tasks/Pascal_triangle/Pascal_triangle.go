package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if j == 0 || i == 0 {
				a[i][j] = 1
			}
		}
	}
	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			a[i][j] = a[i][j-1] + a[i-1][j]
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(strings.Trim(fmt.Sprint(a[i]), "[]"))
	}
}
