package main

import "fmt"

func main() {
	var a, n int
	fmt.Scan(&a, &n)
	pokazatel := 0
	stepen := 1
	for pokazatel < n {
		stepen = stepen * a
		pokazatel++
	}
	fmt.Println(stepen)
}
