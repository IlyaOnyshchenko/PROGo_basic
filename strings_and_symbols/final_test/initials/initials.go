package main

import "fmt"

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)

	fmt.Print(string(a), " ", string(b[0]), ".", string(c[0]), ".")
}
