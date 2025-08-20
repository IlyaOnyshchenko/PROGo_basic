package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)

	r := []rune(c)[0]

	if r >= 97 && r <= 122 {
		fmt.Println(string(r - 32))
	} else {
		if r >= 65 && r <= 90 {
			fmt.Println(string(r + 32))
		}
	}
}
