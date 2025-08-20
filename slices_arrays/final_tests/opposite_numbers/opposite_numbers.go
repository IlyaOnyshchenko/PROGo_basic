package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	var flag bool
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i]+a[j] == 0 && (a[i] < 0 || a[j] < 0) {
				fmt.Println(i, j)
				flag = true
				break
			}
		}
		if flag == true {
			break
		}
	}
}
