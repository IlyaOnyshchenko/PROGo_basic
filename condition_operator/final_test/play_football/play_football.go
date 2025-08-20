package main

import "fmt"

func main() {
	var a int
	var s string
	fmt.Scan(&a, &s)
	if s == "f" {
		fmt.Println("NO")
	} else {
		if a >= 12 && a <= 18 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
