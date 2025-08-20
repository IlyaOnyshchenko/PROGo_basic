package main

import "fmt"

func main() {
	var a, b int
	var c string
	fmt.Scan(&a, &b, &c)
	if c == "/" && b != 0 {
		fmt.Println(float64(a) / float64(b))
	} else {
		if c == "/" && b == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			if c == "*" {
				fmt.Println(a * b)
			} else {
				if c == "+" {
					fmt.Println(a + b)
				} else {
					if c == "-" {
						fmt.Println(a - b)
					} else {
						fmt.Println("Неверная операция")
					}
				}
			}
		}
	}
}
