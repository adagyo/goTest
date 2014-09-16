package utils

import (
	"gopkg.in/mgo.v2"
)

func Connect(conf *Config) (*mgo.Session, int) {
	session, err := mgo.Dial(conf.MgoURL); if err != nil {
		return nil, 1
	}
	return session, 0
}
