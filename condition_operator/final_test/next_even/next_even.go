package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n++
	if n%2 == 0 {
		fmt.Println(n)
	} else {
		n++
		if n%2 == 0 {
			fmt.Println(n)
		}
	}
}
