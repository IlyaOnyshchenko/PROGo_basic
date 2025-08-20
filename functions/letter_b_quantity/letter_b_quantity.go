package main

import (
	"fmt"
)

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	fmt.Println(findB(str1) + findB(str2))
}
func findB(s string) int {
	var bk int
	for i := 0; i < len(s); i++ {
		if s[i] == 98 {
			bk++
		}
	}
	return bk
}
