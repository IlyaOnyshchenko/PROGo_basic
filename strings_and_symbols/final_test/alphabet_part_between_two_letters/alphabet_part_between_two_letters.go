package main

import "fmt"

func main() {
	var c, b string
	fmt.Scan(&c, &b)

	symb1 := []rune(c)[0]
	symb2 := []rune(b)[0]

	if symb1 <= symb2 {
		for letter := symb1; letter <= symb2; letter++ {
			fmt.Print(string(letter), " ")
		}
	} else {
		if symb2 <= symb1 {
			for letter := symb2; letter <= symb1; letter++ {
				fmt.Print(string(letter), " ")
			}
		}
	}
}
