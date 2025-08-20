package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fact := 1
	for i := a; i <= b; i++ {
		fact *= i
	}
	fmt.Println(fact)
}
