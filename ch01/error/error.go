package main

import "errors"

func main() {
	err := doesReturnError()
	if err != nil {
		panic(err)
	}
}

func doesReturnError() error {
	return errors.New("this function simply returns an error")
}
