package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateExposicion : Metodo de insercion de una nueva Exposicion
func CreateExposicion(writter http.ResponseWriter, request *http.Request) {
	var Exposicion models.Exposicion
	
	err := json.NewDecoder(request.Body).Decode(&Exposicion)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var ExposicionValues []interface{}
	var ExposicionStrings []string
	ExposicionValues = utilities.ObjectValues(Exposicion)
	ExposicionStrings = utilities.ObjectFields(Exposicion)

	result, err := utilities.InsertObject("Exposicion", ExposicionValues, ExposicionStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Exposicion creada")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}