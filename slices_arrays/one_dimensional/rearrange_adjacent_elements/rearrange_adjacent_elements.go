package main

import "fmt"

func main() {
	var size, a, b int
	fmt.Scan(&size)
	numbers := make([]int, size)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}
	for i := 0; i < (len(numbers) - 1); i += 2 {
		a = numbers[i]
		b = numbers[i+1]
		numbers[i] = b
		numbers[i+1] = a
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
}
