package main

import "fmt"

func main() {
	var n, a, b, c int
	fmt.Scan(&n)
	a = n % 10
	n = n / 10
	b = n % 10
	n = n / 10
	c = n % 10
	if a > b && b > c && a > c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
