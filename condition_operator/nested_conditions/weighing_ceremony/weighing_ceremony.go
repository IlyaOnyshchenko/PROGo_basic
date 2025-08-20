package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x < 60 {
		fmt.Println("Легкий вес")
	} else {
		if x >= 60 && x < 64 {
			fmt.Println("Первый полусредний вес")
		} else {
			fmt.Println("Полусредний вес")
		}
	}
}
