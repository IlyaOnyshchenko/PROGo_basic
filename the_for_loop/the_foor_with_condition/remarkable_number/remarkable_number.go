package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := n
	sum := 0
	for a > 0 {
		digit := a % 10
		sum += digit
		a /= 10
	}
	if n%sum == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
