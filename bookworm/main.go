package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	bookworms, err := laodBookworms("./testdata/bookworm.json")
	if err != nil {
		fmt.Println("an error occured while loading file")
	}
	for i, value := range bookworms {
		fmt.Println(i)
		fmt.Println(value)
	}
}
