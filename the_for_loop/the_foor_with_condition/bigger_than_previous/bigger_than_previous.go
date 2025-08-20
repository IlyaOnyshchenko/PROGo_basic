package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	digit := 0
	for n != 0 {
		digit = n
		fmt.Scan(&n)
		if n > digit {
			cnt++
		}
	}
	fmt.Println(cnt)
}
