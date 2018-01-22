package main

import(
	"encoding/json"
	"os"
	"log"
	"fmt"
)

type myTown struct {
	TownName string
	Area float32
	Districts []string
}

func main(){
	//Open named file for reading
	file, err := os.Open("resource/tartu.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	//*File satisfies io.Reader interface type so we can
	//use it to define decoder
	decoder := json.NewDecoder(file) // returns decoder *Decoder struct
	//Define emty instance tartu type of myTown
	tartu := myTown{}

	err = decoder.Decode(&tartu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tartu.TownName, tartu.Area, tartu.Districts, tartu.Districts[0])
}

