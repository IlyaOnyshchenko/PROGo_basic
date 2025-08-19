package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	if (a == b && c == d) || (math.Abs(a-c) == math.Abs(b-d)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
