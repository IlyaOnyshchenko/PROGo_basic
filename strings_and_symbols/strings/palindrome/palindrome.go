package main

import (
	"fmt"
)

func main() {
	var str string
	var f bool
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		if string(str[i]) == string(str[len(str)-1-i]) {
			continue
		} else {
			f = true
			fmt.Println("NO")
			break
		}
	}
	if f == false {
		fmt.Println("YES")
	}
}
