package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	for k := 1; k <= n; k++ {
		digit := k
		for digit != 0 {
			digit2 := digit % 10
			if digit2 == 7 {
				cnt++
			}
			digit = digit / 10
		}
	}
	fmt.Println(cnt)
}
