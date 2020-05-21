package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateSelloLibro : Metodo de insercion de un nuevo SelloLibro
func CreateSelloLibro(writter http.ResponseWriter, request *http.Request) {
	var SelloLibro models.SelloLibro
	
	err := json.NewDecoder(request.Body).Decode(&SelloLibro)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var SelloLibroValues []interface{}
	var SelloLibroStrings []string
	SelloLibroValues = utilities.ObjectValues(SelloLibro)
	SelloLibroStrings = utilities.ObjectFields(SelloLibro)

	result, err := utilities.InsertObject("SelloLibro", SelloLibroValues, SelloLibroStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "SelloLibro creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}