package main

import "fmt"

func main() {
	var y, a, b, c, d, e int
	fmt.Scan(&y)
	for i := y; i > 0; i++ {
		e = i + 1
		a = e % 10
		e /= 10
		b = e % 10
		e /= 10
		c = e % 10
		e /= 10
		d = e % 10
		e /= 10
		if a != b && a != c && a != d && b != c && b != d && c != d {
			fmt.Print(d)
			fmt.Print(c)
			fmt.Print(b)
			fmt.Print(a)
			break
		} else {
			continue
		}
	}
}
