package main

import (
	"fmt"
)

func main() {
	var cnt, cnt1 int
	fmt.Scan(&cnt)
	var str string
	fmt.Scan(&str)
	for i := 0; i < cnt-2; i++ {
		if (str[i] == str[i+1]) && (str[i] == str[i+2]) && (str[i+1] == str[i+2]) && (str[i] == 120) {
			cnt1++
		}
	}
	fmt.Println(cnt1)
}
