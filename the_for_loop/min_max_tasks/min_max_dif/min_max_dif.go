package main

import "fmt"

func main() {
	var n, num, max, min int
	_, _ = fmt.Scan(&n, &max)
	min = max
	for i := 1; i <= n; i++ {
		_, _ = fmt.Scan(&num)
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	if max == 0 || min == 0 {
		fmt.Print(0)
	} else {
		fmt.Println(max - min)
	}

}
