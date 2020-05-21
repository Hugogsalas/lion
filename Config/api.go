package main

import (
	"log"
	"net/http"
	"../Controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/User/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/User/get", controllers.GetUser).Methods("POST")
	router.HandleFunc("/Lib/create", controllers.CreateLib).Methods("POST")
	router.HandleFunc("/Lib/get", controllers.GetLib).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
