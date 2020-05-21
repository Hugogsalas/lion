package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateAutor : Metodo de insercion de una  nueva autor
func CreateAutor(writter http.ResponseWriter, request *http.Request) {
	var autor models.Autor
	err := json.NewDecoder(request.Body).Decode(&autor)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var autorValues []interface{}
	var autorStrings []string
	autorValues = utilities.ObjectValues(autor)
	autorStrings = utilities.ObjectFields(autor)

	result, err := utilities.InsertObject("autor", autorValues, autorStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Autor Creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}
