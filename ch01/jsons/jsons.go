package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type PacktBook struct {
	Title          string    `json:'title'`
	Author         string    `json:'author'`
	Length         string    `json:'length'`
	PublishingDate time.Time `json:'pub_date'`
}

func main() {
	book := PacktBook{}
	jsonBook, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Could not marshal object")
	}
	fmt.Println(string(jsonBook))
}
