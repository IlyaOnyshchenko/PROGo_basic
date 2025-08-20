package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&numbers[i])
	}
	min := 10001
	for i := 0; i < count; i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	for i := 0; i < count; i++ {
		numbers[i] = numbers[i] - min
	}
	for i := 0; i < count; i++ {
		fmt.Print(numbers[i], " ")
	}
}
