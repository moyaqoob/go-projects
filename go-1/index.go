package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

type book struct{
	author string `json:"author"`
	title string `json:"title"`
}

type Bookworm struct{
	name string `json:"name"`
    books []book `json:"books"`
}

func main() {
   file:= loadBookworms(`testdata/bookworms.json`)
}

func loadBookworms(filePath string)([]Bookworm,error){
	file,err:= os.Open(filePath)

	if err!=nil{
		panic(err.Error());
	}

	defer file.Close()
	var bookworm []Bookworm
	err = json.NewDecoder(file).Decode(&bookworm)

	if err!=nil{
		return nil,err
	}

	return bookworm,nil

	

	
}