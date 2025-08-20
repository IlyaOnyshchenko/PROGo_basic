package main

import "fmt"

func main() {
	var a, b, max int
	fmt.Scan(&a, &b)
	max = a
	for i := a; i <= b; i++ {
		if i%7 == 0 && i > max {
			max = i
		}
	}
	if max%7 == 0 {
		fmt.Println(max)
	} else {
		fmt.Println("NO")
	}
}
