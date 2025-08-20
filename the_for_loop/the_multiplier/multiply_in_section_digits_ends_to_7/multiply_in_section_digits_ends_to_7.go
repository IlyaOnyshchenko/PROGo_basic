package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fact := 1
	s := 0
	for i := a; i <= b; i++ {
		s = i % 10
		if (s == 7) || (s == -7) {
			fact *= i
		}
	}
	fmt.Println(fact)
}
