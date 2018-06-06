package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type League struct{
	Teams []Footballteam
}

type Footballteam struct {
	Name string    `json:"name"`
	ShortName  string `json:"shortName"`
}


func main() {

	/*Example of a switch and command line arguments
	  together with slice of structure marshalling*/

	//Get the first argument (go run 03.marshaling_slices_of_structs.go PREMIERLEAGUE)
	arg := os.Args[1]

	switch{
	case arg == "":
		fmt.Println("Please provide argument")
	case arg == "PREMIERLEAGUE":

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

	default:
		fmt.Println(arg, " not implemented yet")
	}
}

