package main

import "fmt"

func addN(m int) func(int) int {
	return func(n int) int {
		return n + m
	}
}

func main() {

	addFive := addN(5)
	result := addFive(6)
	fmt.Println(result)
	result = addFive(6)
	fmt.Println(result)
}
