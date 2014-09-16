package entities

import (
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

type User struct {
	Uid      int
	Login    string
	Password string
	Name     string
}

func GetUserByUid(uid int, conf) User {

	usersCollection := db.C("users")

	var u User
	/*err = usersCollection.Find(bson.M{Uid: uid}).One(&u)
	if err != nil {
		fmt.Println("got an error finding user with Uid: " + strconv.Itoa(uid))
	}*/
	return u
}
