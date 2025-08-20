package main

import "fmt"

func main() {
	var size, cnt int
	fmt.Scan(&size)
	numbers := make([]int, size)
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}
	for i := 0; i < (len(numbers) - 1); i++ {
		if numbers[i] < numbers[i+1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
