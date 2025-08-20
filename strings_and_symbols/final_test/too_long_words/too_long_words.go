package main

import (
	"fmt"
)

func main() {
	var cnt int
	var word string
	fmt.Scan(&cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&word)
		if len(word) <= 10 && len(word) > 0 {
			fmt.Println(word)
		} else {
			fmt.Print(string(word[0]), len(word)-2, string(word[len(word)-1]))
			fmt.Println()
		}
	}

}
