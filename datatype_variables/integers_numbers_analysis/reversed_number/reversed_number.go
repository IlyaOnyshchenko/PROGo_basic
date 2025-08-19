package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, n1, n2, n3 int
	fmt.Scan(&n)
	n1 = n % 10
	str1 := strconv.Itoa(n1)
	n = n / 10
	n2 = n % 10
	str2 := strconv.Itoa(n2)
	n = n / 10
	n3 = n % 10
	str3 := strconv.Itoa(n3)
	fmt.Print(str1 + str2 + str3)
}
