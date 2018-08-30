package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type PacktBook struct {
	Title          string    `json:'title'`
	Author         string    `json:'author'`
	Length         int       `json:'length'`
	PublishingDate time.Time `json:'pub_date'`
}

func main() {
	book1 := PacktBook{}
	jsonBook1, err := json.Marshal(book1)
	if err != nil {
		fmt.Println("Could not marshal object")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonBook1))
	lotr := `{"title": "Lord of the Rings", "author": "JRR Tolkien", "pub_date": "1950-01-03 08:00:00 +0000 UTC", "length": 588}`
	jsonBook2 := []byte(lotr)
	var book2 PacktBook
	err = json.Unmarshal(jsonBook2, &book2)
	if err != nil {
		fmt.Println("Could not unmarshal object")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(book2)

}
