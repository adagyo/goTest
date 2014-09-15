package fixtures

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type User struct {
	Uid      int
	Login    string
	Password string
	Name     string
}

func LoadUsers() {
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
		usersCollection.Insert(&User{Uid: i, Login: "user_" + strconv.Itoa(i), Password: "secret", Name: "User #" + strconv.Itoa(i)})
	}

}
