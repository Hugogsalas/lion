package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateSala : Metodo de insercion de una  nueva sala
func CreateSala(writter http.ResponseWriter, request *http.Request) {
	var sala models.Sala
	err := json.NewDecoder(request.Body).Decode(&sala)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var salaValues []interface{}
	var salaStrings []string
	salaValues = utilities.ObjectValues(sala)
	salaStrings = utilities.ObjectFields(sala)

	result, err := utilities.InsertObject("sala", salaValues, salaStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Sala Creada")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

