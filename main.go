package main

import (
	"./entities"
	"./fixtures"
	"./utils"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var conf *utils.Config

func ListUsers(writer http.ResponseWriter, request *http.Request) {
	//requestParams = mux.Vars(request)

	writer.Write([]byte("Je liste les users"))
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	requestParams := mux.Vars(request)

	uid, err := strconv.Atoi(requestParams["id"])
	if err != nil {
		writer.Write([]byte(err))
	}

	writer.Write([]byte("J'affiche le users #" + requestParams["id"]))
	j, _ := json.Marshal(User.GetUserByUid(uid, conf))
	writer.Write([]byte(j))
}

func main() {
	// Load configuration
	utils.LoadConfig(conf)

	// Connect to Mongo and select the database
	sess, ErrNo := utils.Connect(conf)
	if ErrNo > 0 {
		panic("[FATAL] Could not connect to mongo URL '" + conf.MgoURL + "'.")
	}
	defer sess.Close()

	fixtures.LoadUsers(sess, conf)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users", ListUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id:[0-9]*}", GetUser).Methods("GET")

	http.ListenAndServe(":8000", router)
}
