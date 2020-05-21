package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateTipoTaller : Metodo de insercion de un nuevo Tipo Taller
func CreateTipoTaller(writter http.ResponseWriter, request *http.Request) {
	var TipoTaller models.TiposTalleres
	
	err := json.NewDecoder(request.Body).Decode(&TipoTaller)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var TipoTallerValues []interface{}
	var TipoTallerStrings []string
	TipoTallerValues = utilities.ObjectValues(TipoTaller)
	TipoTallerStrings = utilities.ObjectFields(TipoTaller)

	result, err := utilities.InsertObject("TiposTalleres", TipoTallerValues, TipoTallerStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "TipoTaller creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}