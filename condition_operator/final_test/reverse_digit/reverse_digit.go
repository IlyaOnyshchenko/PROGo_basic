package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)
	if x == 0 {
		fmt.Println("Обратного числа не существует")
	} else {
		fmt.Println(math.Pow(x, -1))
	}
}
