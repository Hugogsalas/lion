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
	router.HandleFunc("/Editorial/create", controllers.CreateEditorial).Methods("POST")
	router.HandleFunc("/Exposicion/create", controllers.CreateExposicion).Methods("POST")
	router.HandleFunc("/TipoTaller/create", controllers.CreateTipoTaller).Methods("POST")
	router.HandleFunc("/TipoExposicion/create", controllers.CreateTipoExposicion).Methods("POST")
	router.HandleFunc("/Sello/create", controllers.CreateSello).Methods("POST")
	router.HandleFunc("/Stan/create", controllers.CreateStan).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
