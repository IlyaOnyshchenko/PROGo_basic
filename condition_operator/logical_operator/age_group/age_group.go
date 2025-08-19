package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a <= 13 {
		fmt.Println("детство")
	}
	if 14 <= a && a <= 24 {
		fmt.Println("молодость")
	}
	if 25 <= a && a <= 59 {
		fmt.Println("зрелость")
	}
	if a >= 60 {
		fmt.Println("старость")
	}
}
