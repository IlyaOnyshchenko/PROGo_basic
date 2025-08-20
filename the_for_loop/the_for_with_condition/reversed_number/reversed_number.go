package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var str string
	fmt.Scan(&n)
	for n > 0 {
		digit := n % 10
		str = str + strconv.Itoa(digit)
		n = n / 10
	}
	fmt.Println(str)
}
