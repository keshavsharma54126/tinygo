package main

import (
	"encoding/json"
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

func loadBookworms(filepath string) ([]Bookworms, error) {
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
	return bookworms, nil
}
