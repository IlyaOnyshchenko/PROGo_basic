package main

import (
	"fmt"
)

func main() {
	var n int
	var fd float32 = 8192.0
	fmt.Scan(&n)
	fmt.Println(float32(n) / fd)
}
