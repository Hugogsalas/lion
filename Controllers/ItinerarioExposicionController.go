package controllers

import (
	"encoding/json"
	"net/http"

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