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
			a[i][j] = (j + 1) * (i + 1)
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(strings.Trim(fmt.Sprint(a[i]), "[]"))
	}
}
