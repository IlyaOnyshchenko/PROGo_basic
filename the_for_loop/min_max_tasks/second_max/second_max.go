package main

import "fmt"

func main() {
	var n, max1, max2 int
	fmt.Scan(&n)
	max1 = n
	max2 = 0
	for n != 0 {
		fmt.Scan(&n)
		if n > max1 && n > max2 {
			max2 = max1
			max1 = n
		}
		if n > max2 && n < max1 {
			max2 = n
		}
	}
	fmt.Println(max2)
}
