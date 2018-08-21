package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	myJsonMap := make(map[string]interface{})
	jsonData := []byte(`{"hello": "world"}`)
	err := json.Unmarshal(jsonData, &myJsonMap)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(myJsonMap)
}
