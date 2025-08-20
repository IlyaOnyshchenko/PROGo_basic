package main

import (
	"fmt"
	"strings"
)

func main() {
	var rowsCounts int
	fmt.Scan(&rowsCounts)
	a := make([][]int, rowsCounts)
	for i := range a {
		a[i] = make([]int, rowsCounts)
	}
	for j := len(a) - 1; j >= 1; j-- {
		for i := 1; i < len(a); i++ {
			a[i][j] = 2
		}
	}
	s := len(a)
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			a[i][j] = 0
		}
		s = s - 1
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if i+j == len(a)-1 {
				a[i][j] = 1
			}
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(strings.Trim(fmt.Sprint(a[i]), "[]"))
	}
}
