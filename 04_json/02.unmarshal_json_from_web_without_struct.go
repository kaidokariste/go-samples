package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//For requests we use https://jsonplaceholder.typicode.com/
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	//Read in the response body
	decoder := json.NewDecoder(resp.Body)
	//Because the data is array, allocate slice of maps
	//to avoid creating structs
	var data []map[string]interface{}

	err = decoder.Decode(&data)
	fmt.Println(data)
}
