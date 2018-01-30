package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var databaseName string
var connectionString string

var session *mgo.Session

//DatabaseInit creates a connection to the database
func DatabaseInit(dbName, connectionstring string) {
	databaseName = dbName
	connectionString = connectionstring + dbName

	session, err = mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
}

func setCollection(collection string) *mgo.Collection {
	return session.DB(databaseName).C(collection)
}

//Insert allows users to add generic objects to a collection in the database
func Insert(collection string, object interface{}) bool {
	c := setCollection(collection)
	err := c.Insert(object)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//GetAll returns an array of all objects in a collection
func GetAll(collection string) []interface{} {
	c := setCollection(collection)
	var list []interface{}
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//RemoveAll will empty a collection
func RemoveAll(collection string) bool {
	c := setCollection(collection)
	_, err := c.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//Test tests connection with simple queries
func Test() {
	var result interface{}
	c := setCollection("objects")
	err = c.Find(bson.M{"first": "mitchell"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v", result)
	addPerson()
	addPerson()
	addPerson()

	fmt.Printf("%v", GetAll("people"))
	RemoveAll("people")
}

func addPerson() {
	pers := Person{FirstName: "Mitchell", LastName: "Strong", IsUHN: true, AddedOn: time.Now()}
	Insert("people", pers)
}
