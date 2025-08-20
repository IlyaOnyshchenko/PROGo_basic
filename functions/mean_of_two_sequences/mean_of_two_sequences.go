package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(srZ(a) + srZ(b))
}

func srZ(x float64) float64 {
	var s float64
	for i := 1; i <= int(x); i++ {
		s = s + float64(i)
	}
	return s / x
}
