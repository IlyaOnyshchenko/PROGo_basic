package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Print(factorial(a), " ", factorial(b), " ", factorial(c))
}

func factorial(number int) int {
	fact := 1
	for i := 1; i <= number; i++ {
		if number%2 == 0 && i%2 == 0 {
			fact = fact * i
		} else {
			if number%2 != 0 && i%2 != 0 {
				fact = fact * i
			}
		}
	}
	return fact
}
