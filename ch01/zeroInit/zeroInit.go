package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func divisibleBy(n, divisor int) (bool, error) {

	if divisor == 0 {
		return false, errors.New("A number cannot be divided by zero")
	}

	return (n%divisor == 0), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <number> <divisor>\n", filepath.Base(os.Args[0]))
		os.Exit(100)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	divisor, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	res, err := divisibleBy(number, divisor)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	if res {
		fmt.Printf("%d is divisible by %d\n", number, divisor)
	} else {
		fmt.Printf("%d is NOT divisible by %d\n", number, divisor)
		os.Exit(101)
	}
}
