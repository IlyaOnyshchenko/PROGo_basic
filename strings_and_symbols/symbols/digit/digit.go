package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)
	if input[0] >= '0' && input[0] <= '9' {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
