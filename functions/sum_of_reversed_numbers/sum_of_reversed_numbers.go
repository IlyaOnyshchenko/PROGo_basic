package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(int((chN(a) + chN(b))))
}

func chN(x int) float64 {
	var ns float64
	s := len(strconv.Itoa(x)) - 1
	for i := s; i >= 0; i-- {
		ns = ns + float64(x%10)*math.Pow(10, float64(i))
		x = x / 10
	}
	return ns
}
