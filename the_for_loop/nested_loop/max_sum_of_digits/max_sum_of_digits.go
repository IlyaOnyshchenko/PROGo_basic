package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	maxsum := 0
	maxn := 0
	for i := 1; i <= n; i++ {
		for k := i; k > 0; k /= 10 {
			sum += k % 10
		}
		if sum > maxsum {
			maxsum = sum
			maxn = i
		}
		sum = 0
	}
	fmt.Print(maxn, maxsum)
}
