package main

import "fmt"

func main() {
	var n, a, b, c, d int
	fmt.Scan(&n)
	a = n % 10
	n = n / 10
	b = n % 10
	n = n / 10
	c = n % 10
	n = n / 10
	d = n % 10
	if a == d && b == c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
