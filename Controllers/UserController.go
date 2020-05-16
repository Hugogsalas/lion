package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateUser : Metodo de insercion de un nuevo usuario
func CreateUser(writter http.ResponseWriter, request *http.Request) {
	var usuario models.Usuario
	
	err := json.NewDecoder(request.Body).Decode(&usuario)
	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var userValues []interface{}
	var useStrings []string
	userValues=utilities.ObjectValues(usuario)
	useStrings=utilities.ObjectFields(usuario)
	
	 
	result,err := utilities.InsertObject("Usuarios",userValues,useStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result{
		json.Set("Exito", true)
		json.Set("Message", "Usuario Creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}
