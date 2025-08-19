package main

import (
	"fmt"
)

func main() {
	var n float64
	var a int
	fmt.Scan(&n)
	a = int(n)
	fmt.Println(n - float64(a))
}
