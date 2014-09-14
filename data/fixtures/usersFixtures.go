package usersFixtures

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	id       int
	login    string
	password string
	name     string
}

func Load() {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	db := session.DB("myapi")
	usersCollection := db.C("users")

	nbUsers, err := usersCollection.Count()
	if nbUsers > 0 {
		fmt.Println("Dropping all data from collection 'users'")
		usersCollection.RemoveAll(bson.M{})
	}

	fmt.Println("Populating collection 'users' with 100 documents")
	for i := 0; i < 100; i++ {
		usersCollection.Insert(&user{i, "user_" + string(i), "secret", "User #" + string(i)})
	}
}
