package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		if str[i] == 46 {
			fmt.Print(0)
		} else {
			if (str[i] == 45) && (str[i+1] == 46) {
				fmt.Print(1)
				i++
			} else {
				if (str[i] == 45) && (str[i+1] == 45) {
					fmt.Print(2)
					i++
				}
			}
		}
	}
}
