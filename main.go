package main

import (
	"./data/fixtures"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"encoding/json"
)

type User struct {
	Uid      int
	Login    string
	Password string
	Name     string
}

func GetUserByUid(uid int) User {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	db := session.DB("myapi")
	usersCollection := db.C("users")

	var u User
	err = usersCollection.Find(bson.M{"Uid": uid}).One(&u)
	if err != nil {
		fmt.Println("got an error finding user with Uid: " + strconv.Itoa(uid))
	}
	return u
}

func ListUsers(writer http.ResponseWriter, request *http.Request) {
	//requestParams = mux.Vars(request)

	writer.Write([]byte("Je liste les users"))
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	requestParams := mux.Vars(request)

	uid,err := strconv.Atoi(requestParams["id"])
	if err != nil {
		uid = 0
	}

	writer.Write([]byte("J'affiche le users #" + requestParams["id"]))
	j,err2 := json.Marshal(GetUserByUid(uid))
	if err2 != nil {
		j = []byte("{}")
	}
	writer.Write([]byte(j))
}

func main() {
	fixtures.LoadUsers()

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users", ListUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id:[0-9]*}", GetUser).Methods("GET")
	//router.HandleFunc("/api/v1/users", ListUsers)

	http.ListenAndServe(":8000", router)
}
