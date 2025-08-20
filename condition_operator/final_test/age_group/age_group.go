package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	switch {
	case x <= 3 && x > 0:
		fmt.Println("начинающий")
	case x <= 7 && x >= 4:
		fmt.Println("младший разработчик")
	case x <= 15 && x >= 5:
		fmt.Println("средний разработчик")
	case x >= 16:
		fmt.Println("старший разработчик")
	}
}
