package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&numbers[i])
	}
	indexMax := 0
	indexMin := 0
	min := 10000
	max := numbers[0]
	for i := 0; i < count; i++ {
		if numbers[i] >= max {
			indexMax = i
			max = numbers[i]
		}
	}
	for i := 0; i < count; i++ {
		if numbers[i] < min {
			indexMin = i
			min = numbers[i]
		}
	}
	tmp := numbers[indexMax]
	numbers[indexMax] = numbers[indexMin]
	numbers[indexMin] = tmp
	for i := 0; i < count; i++ {
		fmt.Print(numbers[i], " ")
	}
}
