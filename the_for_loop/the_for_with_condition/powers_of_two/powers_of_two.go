package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	stepen := 1
	for stepen <= n {
		fmt.Println(stepen)
		stepen = stepen * 2
	}
}
