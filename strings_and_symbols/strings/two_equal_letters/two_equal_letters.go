package main

import (
	"fmt"
)

func main() {
	var str string
	var f bool
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if string(str[i]) == string(str[j]) {
				f = true
				fmt.Println(string(str[i]))
				break
			} else {
				continue
			}
		}
		if f == true {
			break
		}
	}
}
