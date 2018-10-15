package main

import "fmt"

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

func main() {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right: &Tree{
				6,
				nil,
				nil,
			},
			Left: nil,
		},
		Left: &Tree{
			LeafValue: 4,
			Right:     nil,
			Left:      nil,
		},
	}

	fmt.Println(root.Right.Right.LeafValue)

}
