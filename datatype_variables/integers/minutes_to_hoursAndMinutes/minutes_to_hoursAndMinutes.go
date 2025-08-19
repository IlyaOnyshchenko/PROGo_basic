package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n, "мин - это", n/60, "час", n%60, "минут.")
}
