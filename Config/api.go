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
	router.HandleFunc("/User/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/Libro/create", controllers.CreateLibro).Methods("POST")
	router.HandleFunc("/Libro/get", controllers.GetLibro).Methods("POST")
	router.HandleFunc("/Editorial/create", controllers.CreateEditorial).Methods("POST")
	router.HandleFunc("/Editorial/get", controllers.GetEditorial).Methods("POST")
	router.HandleFunc("/Exposicion/create", controllers.CreateExposicion).Methods("POST")
	router.HandleFunc("/Exposicion/get", controllers.GetExposicion).Methods("POST")
	router.HandleFunc("/TipoTaller/create", controllers.CreateTipoTaller).Methods("POST")
	router.HandleFunc("/TipoExposicion/create", controllers.CreateTipoExposicion).Methods("POST")
	router.HandleFunc("/Sello/create", controllers.CreateSello).Methods("POST")
	router.HandleFunc("/Sello/get", controllers.GetSello).Methods("POST")
	router.HandleFunc("/Stan/create", controllers.CreateStan).Methods("POST")
	router.HandleFunc("/Sala/create", controllers.CreateSala).Methods("POST")
	router.HandleFunc("/SalaExposicion/create", controllers.CreateSalaExposicion).Methods("POST")
	router.HandleFunc("/SalaTaller/create", controllers.CreateSalaTaller).Methods("POST")
	router.HandleFunc("/Taller/create", controllers.CreateTaller).Methods("POST")
	router.HandleFunc("/Autor/create", controllers.CreateAutor).Methods("POST")
	router.HandleFunc("/Autor/get", controllers.GetAutor).Methods("POST")
	router.HandleFunc("/Itinerario/create", controllers.CreateItinerario).Methods("POST")
	router.HandleFunc("/Itinerario/get", controllers.GetItinerario).Methods("POST")
	router.HandleFunc("/ItinerarioExposicion/create", controllers.CreateItinerarioExposicion).Methods("POST")
	router.HandleFunc("/ItinerarioTaller/create", controllers.CreateItinerarioTaller).Methods("POST")
	router.HandleFunc("/AutorLibro/create", controllers.CreateAutorLibro).Methods("POST")
	router.HandleFunc("/AutorLibro/get", controllers.GetAutorLibro).Methods("POST")
	router.HandleFunc("/EditorialLibro/create", controllers.CreateEditorialLibro).Methods("POST")
	router.HandleFunc("/SelloLibro/create", controllers.CreateSelloLibro).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
