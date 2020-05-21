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
	router.HandleFunc("/Sala/create", controllers.CreateSala).Methods("POST")
	router.HandleFunc("/Autor/create", controllers.CreateAutor).Methods("POST")
	router.HandleFunc("/Itinerario/create", controllers.CreateItinerario).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
