package main

import (
	"github.com/adagyo/goTest/entities"
	"github.com/adagyo/goTest/fixtures"
	"github.com/adagyo/goTest/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"gopkg.in/mgo.v2"
)

var conf utils.Config
var db *mgo.Database
var session *mgo.Session

func ListUsers(writer http.ResponseWriter, request *http.Request) {
	//requestParams = mux.Vars(request)

	writer.Write([]byte("Je liste les users"))
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	requestParams := mux.Vars(request)

	uid, err := strconv.Atoi(requestParams["id"])
	if err != nil {
		writer.Write([]byte("id is not an integer"))
	} else {
		u2find := entities.GetUserByUid(db, uid)
		if u2find.ErrMsg != "" {
			writer.WriteHeader(404)
		} else {
			j, _ := json.Marshal(u2find)
			writer.Write([]byte(j))
		}
	}
}

func main() {
	// Load configuration
	utils.LoadConfig(&conf)

	// Connect to Mongo and select the database
	var ErrNo int
	session, db, ErrNo = utils.Connect(&conf)
	switch ErrNo {
	case 1:
		panic("[FATAL] Could not connect to mongo URL '" + conf.MgoURL + "'.")
	case 2:
		panic("[FATAL] Database '" + conf.MgoDB + "' does not exist.")
	}
	defer session.Close()

	if conf.LoadFixtures == true {
		fixtures.LoadUsers(db)
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users", ListUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id:[0-9]*}", GetUser).Methods("GET")

	http.ListenAndServe(":8000", router)
}
