package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func sum(nums ...string) int {
	res := 0
	for _, x := range nums {
		n, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println(err)
			continue
		}
		res += n
	}
	return res
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: %s <num1> [<num2>...<numN>]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	fmt.Printf("Sum: %d\n", sum(os.Args[1:]...))
}
