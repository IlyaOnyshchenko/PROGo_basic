package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	for n != 1 {
		for i := 2; i <= n; {
			if n%i == 0 {
				fmt.Print(i, " ")
				n = n / i
			} else {
				i++
			}
		}
	}
}
