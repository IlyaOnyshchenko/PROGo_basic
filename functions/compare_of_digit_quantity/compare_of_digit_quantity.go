package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number1, number2 int
	fmt.Scan(&number1, &number2)
	fmt.Println(quaNum(number1, number2))
}
func quaNum(x int, y int) int {
	a := strconv.Itoa(x)
	b := strconv.Itoa(y)
	var s int
	if len(a) == len(b) {
		s = 0
	} else {
		if len(a) < len(b) {
			s = 2
		} else {
			if len(a) > len(b) {
				s = 1
			}
		}
	}
	return s
}
