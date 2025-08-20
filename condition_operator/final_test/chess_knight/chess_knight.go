package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	if (math.Abs(a-c) == 1 && math.Abs(b-d) == 2) || (math.Abs(a-c) == 2 && math.Abs(b-d) == 1) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
