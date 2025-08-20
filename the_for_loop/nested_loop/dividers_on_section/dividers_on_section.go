package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	for k := a; k <= b; k++ {
		cnt := 0
		for i := 1; i <= b; i++ {
			if k%i == 0 {
				cnt++
			}
		}
		if cnt >= c {
			fmt.Print(k, " ")
		}
	}
}
