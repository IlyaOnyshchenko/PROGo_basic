package main

import (
	"fmt"
)

func main() {
	var number1, number2, number3 int
	fmt.Scan(&number1, &number2, &number3)
	fmt.Println(sumSeg(number1, number2) + sumSeg(number2, number3))
}
func sumSeg(a int, b int) int {
	var s int
	if a > b {
		return 0
	}
	for i := a; i <= b; i++ {
		s += i
	}
	return s
}
