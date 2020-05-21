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

//GetEditorial : Metodo que regresa editorial segun parametros
func GetEditorial(writter http.ResponseWriter, request *http.Request) {
	var Editorial models.Editorial
	err := json.NewDecoder(request.Body).Decode(&Editorial)
	jsonResponse := simplejson.New()

	if err == nil {

		var editorialValues []interface{}
		var editorialStrings []string
		editorialValues = utilities.ObjectValues(Editorial)
		editorialStrings = utilities.ObjectFields(Editorial)

		fmt.Println(editorialStrings)
		fmt.Println(editorialValues)
		//Limpia de los atributos del objeto
		if editorialValues[0] == 0 {
			editorialValues[0] = nil
		}

		for i := 1; i < len(editorialStrings); i++ {
			if editorialValues[i] == "" {
				editorialValues[i] = nil
			}
		}

		editorialRows, err := utilities.GetObject("Editorial", nil, editorialStrings, editorialValues)
		if err == nil {
			editorialResultado, err := QueryToEditorial(editorialRows)
			if err == nil {
				if len(editorialResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "La editorial ha sido encontrada")
					jsonResponse.Set("Editorial", editorialResultado)
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

//QueryToEditorial : Metodo que transforma la consulta a objetos Editorial
func QueryToEditorial(result *sql.Rows) ([]models.Editorial, error) {
	var editorialAux models.Editorial
	var recipents []models.Editorial
	for result.Next() {
		err := result.Scan(&editorialAux.ID, &editorialAux.Nombre)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, editorialAux)
	}
	return recipents, nil
}