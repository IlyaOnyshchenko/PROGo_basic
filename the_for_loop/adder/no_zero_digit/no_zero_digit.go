package main

import "fmt"

func main() {
	var x, number int
	fmt.Scan(&x)
	count := 0
	for i := 0; i < x; i++ {
		fmt.Scan(&number)
		if number == 0 {
			count++
		}
	}
	if count > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
