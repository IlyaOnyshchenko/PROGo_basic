package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&numbers[i])
	}
	indexMin := 0
	min := 1000001
	for i := 0; i < count; i++ {
		if numbers[i] < min {
			min = numbers[i]
			indexMin = i
		}
	}
	fmt.Println(indexMin)
}
