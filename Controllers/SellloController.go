package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"fmt"
	
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


//GetSello : Metodo que regresa sello segun parametros
func GetSello(writter http.ResponseWriter, request *http.Request) {
	var Sello models.Sello
	err := json.NewDecoder(request.Body).Decode(&Sello)
	jsonResponse := simplejson.New()

	if err == nil {

		var selloValues []interface{}
		var selloStrings []string
		selloValues = utilities.ObjectValues(Sello)
		selloStrings = utilities.ObjectFields(Sello)

		fmt.Println(selloStrings)
		fmt.Println(selloValues)
		//Limpia de los atributos del objeto
		if selloValues[0] == 0 {
			selloValues[0] = nil
		}

		for i := 1; i < len(selloStrings); i++ {
			if selloValues[i] == "" {
				selloValues[i] = nil
			}
		}

		selloRows, err := utilities.GetObject("Sello", nil, selloStrings, selloValues)
		if err == nil {
			selloResultado, err := QueryToSello(selloRows)
			if err == nil {
				if len(selloResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "El sello ha sido encontrada")
					jsonResponse.Set("Editorial", selloResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron editoriales")
				}
			} else {
				jsonResponse.Set("Exito", false)
				jsonResponse.Set("Message", err.Error())
			}

		} else {
			jsonResponse.Set("Exito", false)
			jsonResponse.Set("Message", err.Error())
		}

	} else {
		jsonResponse.Set("Exito", false)
		jsonResponse.Set("Message", err.Error())
	}

	payload, err := jsonResponse.MarshalJSON()
	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//QueryToSello : Metodo que transforma la consulta a objetos Sello
func QueryToSello(result *sql.Rows) ([]models.Sello, error) {
	var selloAux models.Sello
	var recipents []models.Sello
	for result.Next() {
		err := result.Scan(&selloAux.ID, &selloAux.Descripcion)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, selloAux)
	}
	return recipents, nil
}