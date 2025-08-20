package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	for n > 0 {
		digit := n % 10
		if digit == 4 {
			cnt++
		}
		n = n / 10
	}
	fmt.Println(cnt)
}
