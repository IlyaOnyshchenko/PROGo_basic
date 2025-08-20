package main

import (
	"fmt"
)

func main() {
	var number1, number2 int
	fmt.Scan(&number1, &number2)
	if sumNum(number1) > sumNum(number2) {
		fmt.Println(1)
	} else {
		if sumNum(number1) < sumNum(number2) {
			fmt.Println(2)
		} else {
			if sumNum(number1) == sumNum(number2) {
				fmt.Println(0)
			}
		}
	}
}
func sumNum(x int) int {
	var s int
	for x > 0 {
		s = s + x%10
		x = x / 10
	}
	return s
}
