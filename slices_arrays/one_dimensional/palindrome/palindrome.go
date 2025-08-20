package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)
	array1 := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}
	size2 := size
	for i := 0; i < size; i++ {
		array1[i] = numbers[size2-1]
		size2--
	}
	for i := 0; i < size; i++ {
		if array1[i] != numbers[i] {
			fmt.Println("NO")
			break
		}
		if i == size-1 {
			fmt.Println("YES")
		}
	}
}
