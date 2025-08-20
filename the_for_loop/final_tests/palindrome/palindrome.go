package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, n1, digit, newdigit int
	var str string
	fmt.Scan(&n)
	n1 = n
	for n != 0 {
		digit = n % 10
		n /= 10
		str = str + strconv.Itoa(digit)
	}
	newdigit, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Ошибка")
	}
	if n1 == newdigit {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
