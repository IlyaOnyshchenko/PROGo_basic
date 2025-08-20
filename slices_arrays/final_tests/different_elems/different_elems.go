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
	cnt := 0
	for i := -1000; i <= 1000; i++ {
		cnt = 0
		for _, number := range a {
			if number == i {
				cnt++
			}
		}
		if cnt > 0 {
			s++
		}
	}
	fmt.Println(s)
}
