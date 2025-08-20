package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	pol := 0
	otr := 0
	for n != 0 {
		if n > 0 {
			pol++
		} else {
			otr++
		}
		fmt.Scan(&n)
	}
	fmt.Println(pol - otr)
}
