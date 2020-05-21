package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateSalaExposicion : Metodo de insercion de una nueva relacion Sala-Exposicion
func CreateSalaExposicion(writter http.ResponseWriter, request *http.Request) {
	var SalaExposicion models.SalaExposicion
	
	err := json.NewDecoder(request.Body).Decode(&SalaExposicion)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var SalaExposicionValues []interface{}
	var SalaExposicionStrings []string
	SalaExposicionValues = utilities.ObjectValues(SalaExposicion)
	SalaExposicionStrings = utilities.ObjectFields(SalaExposicion)

	result, err := utilities.InsertObject("SalaExposicion", SalaExposicionValues, SalaExposicionStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Sala-Exposicion creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}