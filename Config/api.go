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
	router.HandleFunc("/Sala/create", controllers.CreateSala).Methods("POST")
	router.HandleFunc("/SalaExposicion/create", controllers.CreateSalaExposicion).Methods("POST")
	router.HandleFunc("/SalaTaller/create", controllers.CreateSalaTaller).Methods("POST")
	router.HandleFunc("/Taller/create", controllers.CreateTaller).Methods("POST")
	router.HandleFunc("/Autor/create", controllers.CreateAutor).Methods("POST")
	router.HandleFunc("/Itinerario/create", controllers.CreateItinerario).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
