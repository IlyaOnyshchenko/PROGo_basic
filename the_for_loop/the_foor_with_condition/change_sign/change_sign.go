package main

import (
	"fmt"
)

func main() {
	var n, digit int
	fmt.Scan(&n)
	cnt := 0
	for n != 0 {
		digit = n
		fmt.Scan(&n)
		if (n > 0 && digit < 0) || (n < 0 && digit > 0) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
