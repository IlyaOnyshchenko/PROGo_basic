package main

import (
	"fmt"
)

func main() {
	var str string
	var f bool
	fmt.Scan(&str)
	if str[0] == 95 || (str[0] >= 97 && str[0] <= 122) || (str[0] >= 65 && str[0] <= 90) {
		f = false
	} else {
		f = true
	}
	for i := 1; i < len(str); i++ {
		if (str[i] == 95) || (str[i] >= 97 && str[i] <= 122) || (str[i] >= 65 && str[i] <= 90) || (str[i] >= 48 && str[i] <= 57) {
			continue
		} else {
			f = true
		}
	}
	if f == false {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
