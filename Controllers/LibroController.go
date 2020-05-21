package controllers

import (
	
	"encoding/json"
	
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

 //CreateLib : Metodo de insercion de un nuevo usuario                                                    
func CreateLib(writter http.ResponseWriter, request *http.Request) {
	var libro models.Libro
	err := json.NewDecoder(request.Body).Decode(&libro)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var libValues []interface{}
	var libStrings []string
	libValues = utilities.ObjectValues(libro)
	libStrings = utilities.ObjectFields(libro)

	result, err := utilities.InsertObject("libro", libValues, libStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Libro añadido")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetLib : Metodo que retorna un libro segun parametros
func GetLib(writter http.ResponseWriter, request *http.Request){
	var params map[string]interface{}
	err := json.NewDecoder(request.Body).Decode(&params)
	json := simplejson.New()
	
	if err!=nil{
		json.Set("Exito",false)
		json.Set("Mensaje",err)
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}