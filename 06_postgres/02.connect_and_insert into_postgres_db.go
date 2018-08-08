package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbhost     = "localhost"
	dbport     = 5432
	dbuser     = "my.username"
	dbpassword = "my.password"
	db         = "mydb"
)

func main() {
	//Create connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpassword, db)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `INSERT INTO myschema.api_response (response_date, response_status, response_body)
VALUES ($1, $2, $3) RETURNING id`

	id := 0
	err = db.QueryRow(sqlStatement, "Thu, 01 Feb 2018 10:14:33 GMT", "200 OK", "Well done!").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
