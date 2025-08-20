package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	for n > 0 || sum > 9 {
		if n == 0 {
			n = sum
			sum = 0
		}
		sum += n % 10
		n /= 10
	}
	fmt.Println(sum)
}
