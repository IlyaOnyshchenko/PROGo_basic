package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, r float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	r = math.Abs(x1-x2) + math.Abs(y1-y2)
	fmt.Println(r)
}
