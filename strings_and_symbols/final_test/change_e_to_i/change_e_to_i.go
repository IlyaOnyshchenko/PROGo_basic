package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)
	for i := 0; i < len(c); i++ {
		if c[i] == 'e' {
			fmt.Print(string('i'))
		} else {
			fmt.Print(string(c[i]))
		}
	}
}
