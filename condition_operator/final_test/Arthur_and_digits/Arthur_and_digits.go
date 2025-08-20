package main

import "fmt"

func main() {
	var k2, k3, k5, k6, s int
	fmt.Scan(&k2, &k3, &k5, &k6)
	s = 256 * min(k2, k5, k6)
	k2 = k2 - min(k2, k5, k6)
	s = s + 32*min(k2, k3)
	fmt.Println(s)
}
