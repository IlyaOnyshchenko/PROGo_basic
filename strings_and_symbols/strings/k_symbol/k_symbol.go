package main

import (
	"fmt"
)

func main() {
	var str string
	var n int
	fmt.Scan(&str)
	fmt.Scan(&n)
	for i := 0; i < len(str); i++ {
		if i != n-1 {
			fmt.Print(string(str[i]))
		}
	}
}
