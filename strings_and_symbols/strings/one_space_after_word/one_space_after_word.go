package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	str := myScanner.Text()
	for i := 1; i <= len(str); i++ {
		if (string(str[i-1]) == " ") && (string(str[i]) == " ") {
			continue
		} else {
			fmt.Print(string(str[i-1]))
		}
	}
}
