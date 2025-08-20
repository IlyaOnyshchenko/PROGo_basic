package main

import (
	"fmt"
)

func main() {
	var input string
	var symb int
	fmt.Scan(&input)
	fmt.Scan(&symb)
	if (symb <= len(input)) && (symb >= 0) {
		fmt.Println(string(input[symb-1]))
	} else {
		fmt.Println("NO")
	}
}
