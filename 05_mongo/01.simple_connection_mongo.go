package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// bson - how it acts when document is inserted
// json - How fields act when json is marshalled
type pet struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	PuppyType string
	Age       int
	NickName  string
	Name      string
	Timestamp time.Time
}

func main() {
	//Create database session
	session, err := mgo.Dial("mongodb://<username>:<password>@ds112233.mlab.com:13098/testpolygon")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//Choose Collection myPets.
	c := session.DB("").C("myPets")

	//Insert new documents to collection
	err = c.Insert(
		&pet{PuppyType: "Rabbit", Age: 4, NickName: "Joe", Name: "Joel II", Timestamp: time.Now()},
		&pet{PuppyType: "Kangaroo", Age: 15, Timestamp: time.Now()},
	)

	if err != nil {
		panic(err)
	}

	//Find one existing document from database
	result := pet{}
	err = c.Find(bson.M{"puppytype": "Rabbit"}).One(&result)
	fmt.Println(result)

	//Find all results
	var results []pet
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
}
