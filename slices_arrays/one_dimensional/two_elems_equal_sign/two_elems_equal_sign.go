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
		if ((numbers[i] > 0) && (numbers[i+1] > 0)) || ((numbers[i] < 0) && (numbers[i+1] < 0)) {
			cnt++
		}
	}
	if cnt > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
