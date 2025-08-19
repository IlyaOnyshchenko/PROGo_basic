package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	text1 := scanner.Text()
	_ = scanner.Scan()
	text2 := scanner.Text()
	_ = scanner.Scan()
	text3 := scanner.Text()
	fmt.Println(text1)
	fmt.Println(text2)
	fmt.Println(text3)
}
