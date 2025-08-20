package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, digit, newdigit int
	var str string
	fmt.Scan(&n)
	for n != 0 {
		digit = n % 10
		n /= 10
		if (digit != 7) && (digit != 5) {
			str = strconv.Itoa(digit) + str
		}
	}
	newdigit, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Ошибка")
	}
	fmt.Println(newdigit)
}
