package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	stepen := 1
	pokazatel := 0
	for stepen < n {
		stepen *= 2
		pokazatel++
	}
	fmt.Println(pokazatel)
}
