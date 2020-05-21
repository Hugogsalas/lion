package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateItinerario : Metodo de insercion de un itinerario
func CreateItinerario(writter http.ResponseWriter, request *http.Request) {
	var itinerario models.Itinerario
	err := json.NewDecoder(request.Body).Decode(&itinerario)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var itinerarioValues []interface{}
	var itinerarioStrings []string
	itinerarioValues = utilities.ObjectValues(itinerario)
	itinerarioStrings = utilities.ObjectFields(itinerario)

	result, err := utilities.InsertObject("itinerario", itinerarioValues, itinerarioStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Itinerario Creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

