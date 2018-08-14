package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <num>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch num {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	default:
		fmt.Printf("Number is %d\n", num)
	}

}
