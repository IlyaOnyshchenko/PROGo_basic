package main

import (
	"fmt"
)

func main() {
	var str string
	var f bool
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "x" || string(str[i]) == "w" {
			f = true
			fmt.Println(string(str[i]))
			break
		}
	}
	if f != true {
		fmt.Println(-1)
	}
}
