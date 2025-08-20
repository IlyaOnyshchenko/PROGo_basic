package main

import "fmt"

func main() {
	var c string
	var s int
	fmt.Scan(&c)
	for i := 0; i < len(c); i++ {
		if c[i] == 'a' || c[i] == 'e' || c[i] == 'i' || c[i] == 'o' || c[i] == 'u' {
			s++
		}
	}
	fmt.Println(s)
}
