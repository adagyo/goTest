package fixtures

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"./../entities"
	"./../utils"
)

func LoadUsers(sess *mgo.Session, conf *utils.Config) {
	usersCollection := sess.DB(conf.MgoDB).C("users")

	nbUsers, _ := usersCollection.Count()
	if nbUsers > 0 {
		fmt.Println("Dropping all data from collection 'users'")
		usersCollection.RemoveAll(bson.M{})
	}

	fmt.Println("Populating collection 'users' with 100 documents")
	for i := 0; i < 100; i++ {
		usersCollection.Insert(&entities.User{Uid: i, Login: "user_" + strconv.Itoa(i), Password: "secret", Name: "User #" + strconv.Itoa(i)})
	}

}
