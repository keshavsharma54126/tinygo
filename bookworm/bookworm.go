package main 

import (
	"os"
	"encoding/json"
)

type Bookworm struct{
	Name string `json:"name"`
	Books []Books `json:"books"`
}

type Books struct{
	Author string `json:"author"`
	Title string `json:"title"`
}

func loadBooworms(filepath string)([]Bookworm,error){
	f,err:= os.Open(filepath)

	if err!=nil{
		return nil,err
	}
	defer f.Close()

	var bookworms []Bookworm

	err=json.NewDecoder(f).Decode(&bookworms)
	if err!=nil{
		return nil,err
	}
	
	return bookworms,nil
}