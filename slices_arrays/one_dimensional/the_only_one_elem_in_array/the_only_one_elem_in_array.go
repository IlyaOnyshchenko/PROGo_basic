package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	numbers := make([]int, count)
	numbers1 := make([]int, 0)
	for i := 0; i < count; i++ {
		fmt.Scan(&numbers[i])
		numbers1 = append(numbers1, numbers[i])
	}
	for i := range numbers {
		cnt := 0
		for j := range numbers1 {
			if numbers1[j] == numbers[i] {
				cnt++
			}
		}
		if cnt == 1 {
			fmt.Print(numbers1[i], " ")
		}
	}
}
