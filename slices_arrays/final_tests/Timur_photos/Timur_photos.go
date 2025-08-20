package main

import (
	"fmt"
)

func main() {
	var n, m int
	var f bool
	fmt.Scan(&n, &m)
	a := make([][]string, n)
	for i := 0; i < n; i++ {
		a[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == "C" || a[i][j] == "M" || a[i][j] == "Y" {
				f = true
				fmt.Println("#Color")
				break
			}
		}
		if f == true {
			break
		}
	}
	if f != true {
		fmt.Println("#Black&White")
	}
}
