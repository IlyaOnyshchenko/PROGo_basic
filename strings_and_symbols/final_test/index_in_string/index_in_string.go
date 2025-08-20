package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	a := myScanner.Text()
	myScanner.Scan()
	b := myScanner.Text()
	if a[0] == b[len(b)-1] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
