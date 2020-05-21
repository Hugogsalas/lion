package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateEditorialLibro : Metodo de insercion de un nuevo EditorialLibro
func CreateEditorialLibro(writter http.ResponseWriter, request *http.Request) {
	var EditorialLibro models.EditorialLibro
	
	err := json.NewDecoder(request.Body).Decode(&EditorialLibro)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var EditorialLibroValues []interface{}
	var EditorialLibroStrings []string
	EditorialLibroValues = utilities.ObjectValues(EditorialLibro)
	EditorialLibroStrings = utilities.ObjectFields(EditorialLibro)

	result, err := utilities.InsertObject("EditorialLibro", EditorialLibroValues, EditorialLibroStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "EditorialLibro creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}