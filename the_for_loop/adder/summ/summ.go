package main

import "fmt"

func main() {
	var x, number int
	fmt.Scan(&x)
	sum := 0
	for i := 0; i < x; i++ {
		fmt.Scan(&number)
		sum = sum + number
	}
	fmt.Println(sum)
}
