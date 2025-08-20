package main

import (
	"fmt"
)

func main() {
	var n, s, s1, s2, s3 int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (a[i]+a[j]) == 4 && (i != j) {
				a[i] = 5
				a[j] = 5
			}
		}
	}
	for i := 0; i < n; i++ {
		if a[i] == 5 {
			s1++
		}
		if a[i] == 4 || a[i] == 3 {
			s++
		}
		if a[i] == 2 {
			s2++
		}
		if a[i] == 1 {
			s3++
		}
	}
	if (s3%2 == 0) && (s3 >= 4) {
		s = s + s3/4
	} else {
		if (s2 > 0) && (s3 > 2) {
			s = s + 2
		} else {
			if ((s2 > 0) && (s3 <= 2)) || (s3 <= 2 && s3 > 0) {
				s = s + 1
			}
		}
	}
	s = s1/2 + s
	fmt.Println(s)
}
