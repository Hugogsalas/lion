package main

import (
	"log"
	"net/http"
	"../Controllers"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/User/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/User/login", controllers.LoginUser).Methods("POST")

	router.HandleFunc("/Libro/create", controllers.CreateLibro).Methods("POST")
	router.HandleFunc("/Libro/get", controllers.GetLibro).Methods("POST")
	router.HandleFunc("/Libro/update", controllers.UpdateLibro).Methods("PUT")
	router.HandleFunc("/Libro/delete", controllers.DeleteLibro).Methods("DELETE")

	router.HandleFunc("/Editorial/create", controllers.CreateEditorial).Methods("POST")
	router.HandleFunc("/Editorial/get", controllers.GetEditorial).Methods("POST")
	router.HandleFunc("/Editorial/update", controllers.UpdateEditorial).Methods("PUT")
	router.HandleFunc("/Editorial/delete", controllers.DeleteEditorial).Methods("DELETE")

	router.HandleFunc("/Exposicion/create", controllers.CreateExposicion).Methods("POST")
	router.HandleFunc("/Exposicion/get", controllers.GetExposicion).Methods("POST")
	router.HandleFunc("/Exposicion/fullGet", controllers.GetExposicionWithType).Methods("POST")
	router.HandleFunc("/Exposicion/update", controllers.UpdateExposicion).Methods("PUT")
	router.HandleFunc("/Exposicion/delete", controllers.DeleteExposicion).Methods("DELETE")

	router.HandleFunc("/TiposTalleres/create",controllers.CreateTiposTalleres).Methods("POST")
	router.HandleFunc("/TiposTalleres/get",controllers.GetTiposTalleres).Methods("POST")
	router.HandleFunc("/TiposTalleres/update", controllers.UpdateTiposTaller).Methods("PUT")
	router.HandleFunc("/TiposTalleres/delete", controllers.DeleteTiposTalleres).Methods("DELETE")

	router.HandleFunc("/TiposExposiciones/create", controllers.CreateTiposExposiciones).Methods("POST")
	router.HandleFunc("/TiposExposiciones/get", controllers.GetTiposExposiciones).Methods("POST")
	router.HandleFunc("/TiposExposiciones/update", controllers.UpdateTiposExposiciones).Methods("PUT")
	router.HandleFunc("/TiposExposiciones/delete", controllers.DeleteTiposExposiciones).Methods("DELETE")

	router.HandleFunc("/Sello/create", controllers.CreateSello).Methods("POST")
	router.HandleFunc("/Sello/get", controllers.GetSello).Methods("POST")
	router.HandleFunc("/Sello/update", controllers.UpdateSello).Methods("PUT")
	router.HandleFunc("/Sello/delete", controllers.DeleteSello).Methods("DELETE")

	router.HandleFunc("/Stan/create", controllers.CreateStan).Methods("POST")
	router.HandleFunc("/Stan/get", controllers.GetStan).Methods("POST")
	router.HandleFunc("/Stan/fullGet", controllers.GetStanWithEditorial).Methods("POST")
	router.HandleFunc("/Stan/update", controllers.UpdateStan).Methods("PUT")
	router.HandleFunc("/Stan/delete", controllers.DeleteStan).Methods("DELETE")

	router.HandleFunc("/Sala/create", controllers.CreateSala).Methods("POST")
	router.HandleFunc("/Sala/get", controllers.GetSala).Methods("POST")
	router.HandleFunc("/Sala/update", controllers.UpdateSala).Methods("PUT")
	router.HandleFunc("/Sala/delete", controllers.DeleteSala).Methods("DELETE")

	router.HandleFunc("/SalaExposicion/create", controllers.CreateSalaExposicion).Methods("POST")
	router.HandleFunc("/SalaExposicion/get", controllers.GetSalaExposicion).Methods("POST")
	router.HandleFunc("/SalaExposicion/update", controllers.UpdateSalaExposicion).Methods("PUT")
	router.HandleFunc("/SalaExposicion/delete", controllers.DeleteSalaExposicion).Methods("DELETE")

	router.HandleFunc("/SalaTaller/create", controllers.CreateSalaTaller).Methods("POST")
	router.HandleFunc("/SalaTaller/get", controllers.GetSalaTaller).Methods("POST")
	router.HandleFunc("/SalaTaller/update", controllers.UpdateSalaTaller).Methods("PUT")
	router.HandleFunc("/SalaTaller/delete", controllers.DeleteSalaTaller).Methods("DELETE")

	router.HandleFunc("/Taller/create", controllers.CreateTaller).Methods("POST")
	router.HandleFunc("/Taller/get", controllers.GetTaller).Methods("POST")
	router.HandleFunc("/Taller/update", controllers.UpdateTaller).Methods("PUT")
	router.HandleFunc("/Taller/delete", controllers.DeleteTaller).Methods("DELETE")

	router.HandleFunc("/Autor/create", controllers.CreateAutor).Methods("POST")
	router.HandleFunc("/Autor/get", controllers.GetAutor).Methods("POST")
	router.HandleFunc("/Autor/update", controllers.UpdateAutor).Methods("PUT")
	router.HandleFunc("/Autor/delete", controllers.DeleteAutor).Methods("DELETE")

	router.HandleFunc("/Itinerario/create", controllers.CreateItinerario).Methods("POST")
	router.HandleFunc("/Itinerario/get", controllers.GetItinerario).Methods("POST")
	router.HandleFunc("/Itinerario/update", controllers.UpdateItinerario).Methods("PUT")
	router.HandleFunc("/Itinerario/delete", controllers.DeleteItinerario).Methods("DELETE")

	router.HandleFunc("/ItinerarioExposicion/create", controllers.CreateItinerarioExposicion).Methods("POST")
	router.HandleFunc("/ItinerarioExposicion/get", controllers.GetItinerarioExposicion).Methods("POST")
	router.HandleFunc("/ItinerarioExposicion/update", controllers.UpdateItinerarioExposicion).Methods("PUT")
	router.HandleFunc("/ItinerarioExposicion/delete", controllers.DeleteItinerarioExposicion).Methods("DELETE")

	router.HandleFunc("/ItinerarioTaller/create", controllers.CreateItinerarioTaller).Methods("POST")
	router.HandleFunc("/ItinerarioTaller/get", controllers.GetItinerarioTaller).Methods("POST")
	router.HandleFunc("/ItinerarioTaller/update", controllers.UpdateItinerarioTaller).Methods("PUT")
	router.HandleFunc("/ItinerarioTaller/delete", controllers.DeleteItinerarioTaller).Methods("DELETE")

	router.HandleFunc("/AutorLibro/create", controllers.CreateAutorLibro).Methods("POST")
	router.HandleFunc("/AutorLibro/get", controllers.GetAutorLibro).Methods("POST")
	router.HandleFunc("/AutorLibro/update", controllers.UpdateAutorLibro).Methods("PUT")
	router.HandleFunc("/AutorLibro/delete", controllers.DeleteAutorLibro).Methods("DELETE")

	router.HandleFunc("/EditorialLibro/create", controllers.CreateEditorialLibro).Methods("POST")
	router.HandleFunc("/EditorialLibro/get", controllers.GetEditorialLibro).Methods("POST")
	router.HandleFunc("/EditorialLibro/update", controllers.UpdateEditorialLibro).Methods("PUT")
	router.HandleFunc("/EditorialLibro/delete", controllers.DeleteEditorialLibro).Methods("DELETE")

	router.HandleFunc("/SelloLibro/create", controllers.CreateSelloLibro).Methods("POST")
	router.HandleFunc("/SelloLibro/get", controllers.GetSelloLibro).Methods("POST")
	router.HandleFunc("/SelloLibro/update", controllers.UpdateSelloLibro).Methods("PUT")
	router.HandleFunc("/SelloLibro/delete", controllers.DeleteSelloLibro).Methods("DELETE")

	log.Fatal(http.ListenAndServe(
		":8080", 
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
			handlers.AllowedOrigins([]string{"*"}))(router)))
}
