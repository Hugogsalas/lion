package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateEditorial : Metodo de insercion de una nueva editorial
func CreateEditorial(writter http.ResponseWriter, request *http.Request) {
	var Editorial models.Editorial
	
	err := json.NewDecoder(request.Body).Decode(&Editorial)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var editorialValues []interface{}
	var editorialStrings []string
	editorialValues = utilities.ObjectValues(Editorial)
	editorialStrings = utilities.ObjectFields(Editorial)

	result, err := utilities.InsertObject("Editorial", editorialValues, editorialStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Editorial creada")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}