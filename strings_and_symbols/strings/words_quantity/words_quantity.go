package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s int
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	str := myScanner.Text()
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			s++
		}
	}
	fmt.Println(s + 1)
}
