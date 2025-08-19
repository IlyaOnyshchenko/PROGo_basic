package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if 100 <= a && a <= 999 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
