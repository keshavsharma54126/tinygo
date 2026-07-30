package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Bookworms struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `jsong:"title"`
}

func laodBookworms(filepath string) ([]Bookworms, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	var bookworms []Bookworms
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&bookworms)
	if err != nil {
		return nil, err
	}
	fmt.Println(bookworms)
	return bookworms, nil
}
