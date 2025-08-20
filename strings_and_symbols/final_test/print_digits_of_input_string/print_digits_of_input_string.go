package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	c := myScanner.Text()
	for i := 0; i < len(c); i++ {
		if c[i] >= '0' && c[i] <= '9' {
			fmt.Print(string(c[i]), " ")
		}
	}
}
