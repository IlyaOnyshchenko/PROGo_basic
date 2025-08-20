package main

import (
	"fmt"
)

func main() {
	var rowsCounts int
	fmt.Scan(&rowsCounts)
	a := make([][]int, rowsCounts)
	var f bool
	for i := range a {
		a[i] = make([]int, rowsCounts)
		for j := 0; j < rowsCounts; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	b := make([]int, 1)
	c := make([]int, 1)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if i < j {
				b = append(b, a[i][j])
			}
		}
	}
	s := 0
	for j := 0; j < len(a); j++ {
		s++
		for i := s; i < len(a); i++ {
			if i > j {
				c = append(c, a[i][j])
			}
		}
	}
	for i := 0; i < len(b); i++ {
		if b[i] == c[i] {
			f = true
		} else {
			f = false
			fmt.Println("NO")
			break
		}
	}
	if f == true {
		fmt.Println("YES")
	}
}
