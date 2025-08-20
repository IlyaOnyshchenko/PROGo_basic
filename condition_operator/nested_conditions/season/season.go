package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch {
	case n == 12 || (n <= 2 && n > 0):
		fmt.Println("Зима")
	case 3 <= n && n <= 5:
		fmt.Println("Весна")
	case 6 <= n && n <= 8:
		fmt.Println("Лето")
	case 9 <= n && n <= 11:
		fmt.Println("Осень")
	}
}
