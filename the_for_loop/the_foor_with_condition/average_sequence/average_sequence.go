package main

import (
	"fmt"
)

func main() {
	var n, sum, cnt float64
	fmt.Scan(&n)
	for n != 0 {
		cnt++
		sum = sum + n
		fmt.Scan(&n)
	}
	fmt.Println(float64(sum / cnt))
}
