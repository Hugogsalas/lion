package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateTaller : Metodo de insercion de un nuevo Taller
func CreateTaller(writter http.ResponseWriter, request *http.Request) {
	var Taller models.Taller
	
	err := json.NewDecoder(request.Body).Decode(&Taller)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var TallerValues []interface{}
	var TallerStrings []string
	TallerValues = utilities.ObjectValues(Taller)
	TallerStrings = utilities.ObjectFields(Taller)

	result, err := utilities.InsertObject("Taller", TallerValues, TallerStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Taller creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}