package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	for n > 0 {
		digit := n % 2
		if digit == 1 {
			cnt++
		}
		n = n / 2
	}
	fmt.Println(cnt)
}
