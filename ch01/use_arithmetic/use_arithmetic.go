package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lekum/go-design-patterns/ch01/arithmetic"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <num1> <num2> ...\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	var args = make([]int, len(os.Args[1:]))
	for idx, v := range os.Args[1:] {
		args[idx], _ = strconv.Atoi(v)
	}
	fmt.Println("Sum:", arithmetic.Sum(args...))
	fmt.Println("Substract:", arithmetic.Substract(args...))
	fmt.Println("Multiply:", arithmetic.Multiply(args...))
}
