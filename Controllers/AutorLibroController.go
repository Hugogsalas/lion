package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateAutorLibro : Metodo de insercion de un nuevo AutorLibro
func CreateAutorLibro(writter http.ResponseWriter, request *http.Request) {
	var AutorLibro models.AutorLibro
	
	err := json.NewDecoder(request.Body).Decode(&AutorLibro)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var AutorLibroValues []interface{}
	var AutorLibroStrings []string
	AutorLibroValues = utilities.ObjectValues(AutorLibro)
	AutorLibroStrings = utilities.ObjectFields(AutorLibro)

	result, err := utilities.InsertObject("AutorLibro", AutorLibroValues, AutorLibroStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "AutorLibro creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}
