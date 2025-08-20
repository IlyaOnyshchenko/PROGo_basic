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
		digit := n % 2
		str = str + strconv.Itoa(digit)
		n = n / 2
	}
	fmt.Println(str)
}
