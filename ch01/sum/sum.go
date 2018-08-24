package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func sum(a, b int) int {
	return a + b
}
func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <num1> <num2>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	result := sum(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, result)
}
