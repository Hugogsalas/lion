package controllers

import (
	"encoding/json"
	"net/http"
	"database/sql"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateItinerarioExposicion : Metodo de insercion de un itinerarioExposicion
func CreateItinerarioExposicion(writter http.ResponseWriter, request *http.Request) {
	var itinerarioExposicion models.ItinerarioExposicion
	err := json.NewDecoder(request.Body).Decode(&itinerarioExposicion)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var itinerarioExposicionValues []interface{}
	var itinerarioExposicionStrings []string
	itinerarioExposicionValues = utilities.ObjectValues(itinerarioExposicion)
	itinerarioExposicionStrings = utilities.ObjectFields(itinerarioExposicion)

	result, err := utilities.InsertObject("itinerarioExposicion", itinerarioExposicionValues, itinerarioExposicionStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "ItinerarioExposicion Creada")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetItinerarioExposicion : metodo que retorna una relacion Itinerario-Exposicion
func GetItinerarioExposicion(writter http.ResponseWriter, request *http.Request) {
	var ItinerarioExposicion models.ItinerarioExposicion
	err := json.NewDecoder(request.Body).Decode(&ItinerarioExposicion)
	jsonResponse := simplejson.New()

	if err == nil {
		var valuesHorario interface{} = nil
		if ItinerarioExposicion.Horario != "" {
			valuesHorario = ItinerarioExposicion.Horario
		}
		ItinerarioExposicionRows, err := utilities.CallStorageProcedure("PAItinerarioExposicion", []interface{}{ItinerarioExposicion.IDItinerario, ItinerarioExposicion.IDExposicion, valuesHorario})
		if err == nil {
			var ItinerarioExposicionResultado []map[string]interface{}

			if ItinerarioExposicion.IDItinerario == 0 && ItinerarioExposicion.IDItinerario == 0 {
				if ItinerarioExposicion.Horario!="" {
					ItinerarioExposicionResultado,err=HorariowithActvities(ItinerarioExposicionRows)
				}
			}

			if err == nil {
				if len(ItinerarioExposicionResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "AutorLibro encontrado")
					jsonResponse.Set("ItinerarioExposicion", ItinerarioExposicionResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron Itinerario-Exposicion")
				}
			} else {
				jsonResponse.Set("Exito", false)
				jsonResponse.Set("Message", err.Error())
			}

		} else {
			jsonResponse.Set("Exito", false)
			jsonResponse.Set("Message", err.Error())
		}

	} else {
		jsonResponse.Set("Exito", false)
		jsonResponse.Set("Message", err.Error())
	}

	payload, err := jsonResponse.MarshalJSON()
	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//HorariowithActvities : metodo que retorna una relacion Itinerario-Exposicion
func HorariowithActvities(result *sql.Rows) ([]map[string]interface{},error){
	var response []map[string]interface{}
	
	return response, nil
}