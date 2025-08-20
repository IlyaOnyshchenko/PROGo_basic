package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
		if numbers[i]%3 == 0 {
			fmt.Print(numbers[i], " ")
		}
	}
}
