package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateSalaTaller : Metodo de insercion de una nueva relacion Sala-Taller
func CreateSalaTaller(writter http.ResponseWriter, request *http.Request) {
	var SalaTaller models.SalaTaller
	
	err := json.NewDecoder(request.Body).Decode(&SalaTaller)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var SalaTallerValues []interface{}
	var SalaTallerStrings []string
	SalaTallerValues = utilities.ObjectValues(SalaTaller)
	SalaTallerStrings = utilities.ObjectFields(SalaTaller)

	result, err := utilities.InsertObject("SalaTaller", SalaTallerValues, SalaTallerStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Sala-Taller creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}