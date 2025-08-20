package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	switch {
	case x%2 != 0:
		fmt.Println("YES")
	case (x%2 == 0) && (x >= 2 && x <= 5):
		fmt.Println("NO")
	case (x%2 == 0) && (x >= 6 && x <= 20):
		fmt.Println("YES")
	case (x%2 == 0) && (x > 20):
		fmt.Println("NO")
	}
}
