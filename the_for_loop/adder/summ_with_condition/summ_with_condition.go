package main

import "fmt"

func main() {
	var x, number int
	fmt.Scan(&x)
	sum := 0
	for i := 0; i < x; i++ {
		fmt.Scan(&number)
		if (number%2 == 0) && (number%3 != 0) {
			sum += number
		}
	}
	fmt.Println(sum)
}
