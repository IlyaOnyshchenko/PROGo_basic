package main

import "fmt"

func main() {
	var n, cnt, digit int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&digit)
		if digit == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
