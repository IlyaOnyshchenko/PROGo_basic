package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)

	r := []rune(c)[0]

	for letter := r; letter <= 'z'; letter++ {
		fmt.Print(string(letter), " ")
	}
}
