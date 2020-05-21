package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateSello : Metodo de insercion de un nuevo sello
func CreateSello(writter http.ResponseWriter, request *http.Request) {
	var Sello models.Sello
	
	err := json.NewDecoder(request.Body).Decode(&Sello)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var SelloValues []interface{}
	var SelloStrings []string
	SelloValues = utilities.ObjectValues(Sello)
	SelloStrings = utilities.ObjectFields(Sello)

	result, err := utilities.InsertObject("Sello", SelloValues, SelloStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Sello creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}