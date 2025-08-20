package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := 0
	for i := 1; i <= n; i++ {
		if n%3 == 0 {
			s++
			n /= 3
		}
	}
	fmt.Println(s)
}
