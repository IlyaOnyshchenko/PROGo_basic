package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}
	numbers2 := make([]int, size)
	numbers2[0] = numbers[len(numbers)-1]
	for i := 1; i < len(numbers); i++ {
		numbers2[i] = numbers[i-1]
	}
	for i := 0; i < len(numbers2); i++ {
		fmt.Print(numbers2[i], " ")
	}
}
