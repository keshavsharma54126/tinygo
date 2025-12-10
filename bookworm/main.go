package main

import (
	"fmt"
	"os"
)


func main(){
	bookworms,err:= loadBooworms("testdata/bookworms.json")
	if err!=nil{
		fmt.Fprintf(os.Stderr,"failed to load bookworms:%s\n",err)
		os.Exit(1)
	}

	fmt.Println(bookworms)

}