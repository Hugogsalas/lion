package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateTipoExposicion : Metodo de insercion de un nuevo Tipo Exposiciones
func CreateTipoExposicion(writter http.ResponseWriter, request *http.Request) {
	var TipoExposicion models.TiposExposicion
	
	err := json.NewDecoder(request.Body).Decode(&TipoExposicion)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var TipoExposicionValues []interface{}
	var TipoExposicionStrings []string
	TipoExposicionValues = utilities.ObjectValues(TipoExposicion)
	TipoExposicionStrings = utilities.ObjectFields(TipoExposicion)

	result, err := utilities.InsertObject("TiposExposicion", TipoExposicionValues, TipoExposicionStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "TipoExposicion creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}