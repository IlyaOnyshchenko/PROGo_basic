package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch {
	case n == 1 || n == 3 || n == 5 || n == 7 || n == 8 || n == 10 || n == 12:
		fmt.Println(31)
	case n == 4 || n == 6 || n == 9 || n == 11:
		fmt.Println(30)
	case n == 2:
		fmt.Println(29)
	}
}
