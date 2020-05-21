package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateStan : Metodo de insercion de un nuevo Stan
func CreateStan(writter http.ResponseWriter, request *http.Request) {
	var Stan models.Stan
	
	err := json.NewDecoder(request.Body).Decode(&Stan)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var StanValues []interface{}
	var StanStrings []string
	StanValues = utilities.ObjectValues(Stan)
	StanStrings = utilities.ObjectFields(Stan)

	result, err := utilities.InsertObject("Stan", StanValues, StanStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Stan creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}