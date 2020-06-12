package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

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

	editorialStrings, editorialValues := utilities.ObjectFields(Editorial)

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

		editorialStrings, editorialValues := utilities.ObjectFields(Editorial)

		//Limpia de los atributos del objeto
		if editorialValues[0] == 0 {
			editorialValues[0] = nil
		}

		for i := 1; i < len(editorialStrings); i++ {
			if editorialValues[i] == "" {
				editorialValues[i] = nil
			}
		}

		var editorialQuery models.GetQuery

		editorialQuery.Tables = []string{"Editorial"}
		editorialQuery.Selects = nil
		editorialQuery.Params = [][]string{editorialStrings}
		editorialQuery.Values = [][]interface{}{editorialValues}
		editorialQuery.Conditions = nil

		editorialRows, err := utilities.GetObject(editorialQuery)
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

//UpdateEditorial : Metodo que actualiza editoriales segun parametros
func UpdateEditorial(writter http.ResponseWriter, request *http.Request) {
	var editorial models.Editorial
	err := json.NewDecoder(request.Body).Decode(&editorial)
	jsonResponse := simplejson.New()
	if err == nil {

		var editorialFilters []string
		var editorialFiltersValues []interface{}

		editorialFilters = append(editorialFilters, "ID")
		editorialFiltersValues = append(editorialFiltersValues, editorial.ID)

		editorial.ID=0

		editorialStrings, editorialValues := utilities.ObjectFields(editorial)

		editorialRows, err := utilities.UpdateObject("Editorial", editorialFilters, editorialFiltersValues, editorialStrings, editorialValues)
		if err == nil {

			if editorialRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "Editorial actualizada")

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

//EditorialesToInterfaces : metodo que transforma un arreglo de Editoriales en interfaces
func EditorialesToInterfaces(Editoriales []models.Editorial) []interface{} {
	var arrayInterface []interface{}
	for i := 0; i < len(Editoriales); i++ {
		var editorialInterface interface{}
		editorialInterface = Editoriales[i]
		arrayInterface = append(arrayInterface, editorialInterface)
	}
	return arrayInterface
}
