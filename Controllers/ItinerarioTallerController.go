package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateItinerarioTaller : Metodo de insercion de un nuevo itinerarioTaller
func CreateItinerarioTaller(writter http.ResponseWriter, request *http.Request) {
	var itinerarioTaller models.ItinerarioTaller
	err := json.NewDecoder(request.Body).Decode(&itinerarioTaller)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var itinerarioTallerValues []interface{}
	var itinerarioTallerStrings []string
	itinerarioTallerValues = utilities.ObjectValues(itinerarioTaller)
	itinerarioTallerStrings = utilities.ObjectFields(itinerarioTaller)

	result, err := utilities.InsertObject("ItinerarioTaller", itinerarioTallerValues, itinerarioTallerStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "ItinerarioTaller Creada")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

