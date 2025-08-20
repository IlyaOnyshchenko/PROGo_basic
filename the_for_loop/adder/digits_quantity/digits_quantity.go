package main

import "fmt"

func main() {
	var x, number int
	fmt.Scan(&x)
	count := 0
	for i := 0; i < x; i++ {
		fmt.Scan(&number)
		number = number % 10
		if number == 0 {
			count++
		}
	}
	fmt.Println(count)
}
