package entities

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	UserId      int
	Login    string
	Password string
	Name     string
	ErrMsg   string
}

func GetUserByUid(db *mgo.Database, uid int) User {
	usersCollection := db.C("users")

	u2find := User{}
	errmsg := usersCollection.Find(bson.M{"userid" : uid}).One(&u2find); if errmsg != nil {
		u2find.ErrMsg = errmsg.Error()
	}

	return u2find
}
