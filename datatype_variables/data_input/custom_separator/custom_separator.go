package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	razd := scanner.Text()
	_ = scanner.Scan()
	elem1 := scanner.Text()
	_ = scanner.Scan()
	elem2 := scanner.Text()
	_ = scanner.Scan()
	elem3 := scanner.Text()
	fmt.Print(elem1, razd, elem2, razd, elem3)
}
