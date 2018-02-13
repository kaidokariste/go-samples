package main

import (
	"encoding/json"
	"fmt"
)

type League struct{
	Teams []Footballteam
}

type Footballteam struct {
	Name string    `json:"name"`
	ShortName  string `json:"shortName"`
}


func main() {

	//Define slice of structs conatining two documents
	list := []Footballteam{
		{"Manchester United","MANU"},
		{"Chelsea","CHE"},
	}

	//Marshal them to  Parent struct into Teams variable
	jsonString, err := json.Marshal(League{Teams: list})
	if err != nil{
		panic(err)
	}

	fmt.Println(list)
	fmt.Println(string(jsonString))
}

