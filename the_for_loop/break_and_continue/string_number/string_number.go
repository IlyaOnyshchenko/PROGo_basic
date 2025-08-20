package main

import (
	"fmt"
)

func main() {
	var s, s1 string
	fmt.Scan(&s)
	cnt := 0
	for cnt >= 0 {
		fmt.Scan(&s1)
		cnt++
		if s1 == s {
			break
		} else {
			continue
		}
	}
	fmt.Println(cnt)
}
