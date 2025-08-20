package main

import "fmt"

func main() {
	var count, count2 int
	fmt.Scan(&count)
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&numbers[i])
	}
	fmt.Scan(&count2)
	numbers2 := make([]int, count2)
	for i := 0; i < count2; i++ {
		fmt.Scan(&numbers2[i])
	}
	fmt.Println((findMax(numbers, count)) * (findMax(numbers2, count2)))
}
func findMax(array []int, y int) int {
	max := array[0]
	for i := 1; i < y; i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}
