package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
}
