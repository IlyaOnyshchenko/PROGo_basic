package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&numbers[i])
	}
	flag := "NO"
	for i := -10000; i < 10000; i++ {
		cnt := 0
		for _, number := range numbers {
			if number == i {
				cnt++
			}
		}
		if cnt > 1 {
			flag = "YES"
			fmt.Println(flag)
			break
		}
	}
	if flag != "YES" {
		fmt.Println(flag)
	}
}
