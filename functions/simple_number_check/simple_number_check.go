package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(prCom(number))
}
func prCom(n int) string {
	var str string
	var k int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			k++
		}
	}
	if k == 2 {
		str = "prime"
	} else if k > 2 {
		str = "composite"
	}
	return str
}
