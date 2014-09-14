package main

import (
	"./data/fixtures"
	"github.com/gorilla/mux"
	"net/http"
)

func ListUsers(writer http.ResponseWriter, request *http.Request) {
	//requestParams = mux.Vars(request)

	writer.Write([]byte("Je liste les users"))
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	requestParams := mux.Vars(request)

	writer.Write([]byte("J'affiche le users #" + requestParams["id"]))
}

func main() {
	usersFixtures.Load()

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users", ListUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id:[0-9]*}", GetUser).Methods("GET")
	//router.HandleFunc("/api/v1/users", ListUsers)

	http.ListenAndServe(":8000", router)
}
