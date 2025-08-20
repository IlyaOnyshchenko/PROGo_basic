package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, x1, x2 float64
	fmt.Scan(&a, &b, &c)
	d = math.Pow(b, 2) - 4*a*c
	if d > 0 && a != 0 {
		x1 = (-b + math.Sqrt(d)) / (2 * a)
		x2 = (-b - math.Sqrt(d)) / (2 * a)
		if x1 > x2 {
			fmt.Println(x2)
			fmt.Println(x1)
		} else {
			fmt.Println(x1)
			fmt.Println(x2)
		}
	} else {
		if d == 0 && a != 0 {
			x1 = -b / (2 * a)
			fmt.Println(x1)
		}
	}
}
