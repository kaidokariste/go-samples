package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"encoding/json"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabse"
)

type (
	DataResource struct {
		Data []Questionary `json:"data"`
	}

	Questionary struct {
		Templateid string          `json:"templateid"`
		Email      string          `json:"email"`
		Tags       json.RawMessage `json:"tags"`
	}
)

func main() {

	//Define items instance where we collect set of Questionary structures
	var items []Questionary

	//Create connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Prepare questionary questions
	QSendQuestions := `select templateid,email,tags from myschema.myquestionary`
	rows, err := db.Query(QSendQuestions)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//Loop trough the Query result
	for rows.Next() {
		var (
			templateid, email string
			tags              json.RawMessage
		)

		err = rows.Scan(&templateid, &email, &tags)
		if err != nil {
			panic(err)
		}

		//Compose questionary instance
		item := Questionary{
			templateid,
			email,
			tags,
		}

		//Add them to slice of structs
		items = append(items, item)
	}

	//Marshal them to parent struct into Data variable to get array of json objects
	jsonString, err := json.Marshal(DataResource{Data: items})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonString))
}
