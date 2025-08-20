package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)

	r := []rune(c)[0]

	for letter := 'a'; letter <= r; letter++ {
		fmt.Print(string(letter), " ")
	}
}
