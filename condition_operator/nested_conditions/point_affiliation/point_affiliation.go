package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	if (x <= 1 && x >= -3) || (x <= 9 && x >= 5) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
