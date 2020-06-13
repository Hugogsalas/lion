package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateAutor : Metodo de insercion de un nuevo autor
func CreateAutor(writter http.ResponseWriter, request *http.Request) {
	var autor models.Autor
	err := json.NewDecoder(request.Body).Decode(&autor)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	autorStrings, autorValues := utilities.ObjectFields(autor)

	result, err := utilities.InsertObject("autor", autorValues, autorStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result!=0 {
		json.Set("Exito", true)
		json.Set("Message", "Autor Creado")
		json.Set("Id", result)
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetAutor : Metodo que regresa autores segun parametros
func GetAutor(writter http.ResponseWriter, request *http.Request) {
	var autor models.Autor
	err := json.NewDecoder(request.Body).Decode(&autor)
	jsonResponse := simplejson.New()
	if err == nil {

		autorStrings, autorValues := utilities.ObjectFields(autor)

		var autorQuery models.GetQuery

		autorQuery.Tables = []string{"Autor"}
		autorQuery.Selects = nil
		autorQuery.Params = [][]string{autorStrings}
		autorQuery.Values = [][]interface{}{autorValues}
		autorQuery.Conditions = nil

		autorRows, err := utilities.GetObject(autorQuery)
		if err == nil {
			autoresResultado, err := QueryToAutor(autorRows)
			if err == nil {
				if len(autoresResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "Autores encontrados")
					jsonResponse.Set("Autores", autoresResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron autores")
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

//UpdateAutor : Metodo que actualiza autores segun parametros
func UpdateAutor(writter http.ResponseWriter, request *http.Request) {
	var autor models.Autor
	err := json.NewDecoder(request.Body).Decode(&autor)
	jsonResponse := simplejson.New()
	if err == nil {

		var autorFilters []string
		var autorFiltersValues []interface{}

		autorFilters = append(autorFilters, "ID")
		autorFiltersValues = append(autorFiltersValues, autor.ID)

		autor.ID = 0

		autorStrings, autorValues := utilities.ObjectFields(autor)

		autorRows, err := utilities.UpdateObject("Autor", autorFilters, autorFiltersValues, autorStrings, autorValues)
		if err == nil {

			if autorRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "Autor actualizado")

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

//DeleteAutor : Metodo que elimina autores segun parametros
func DeleteAutor(writter http.ResponseWriter, request *http.Request) {
	var Autor models.Autor
	err := json.NewDecoder(request.Body).Decode(&Autor)
	jsonResponse := simplejson.New()
	if err == nil {

		AutorStrings, AutorValues := utilities.ObjectFields(Autor)

		AutorDel, err := utilities.DeleteObject("Autor", AutorStrings, AutorValues)
		if err == nil {

			if AutorDel {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "Autor eliminado")

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

//QueryToAutor : Metodo que transforma la consulta a objetos Autor
func QueryToAutor(result *sql.Rows) ([]models.Autor, error) {
	var autorAux models.Autor
	var recipents []models.Autor
	for result.Next() {
		err := result.Scan(&autorAux.ID, &autorAux.Nombre, &autorAux.ApellidoPaterno, &autorAux.ApellidoMaterno)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, autorAux)
	}
	return recipents, nil
}

//AutoresToInterfaces : metodo que transforma un arreglo de Autores en interfaces
func AutoresToInterfaces(Autores []models.Autor) []interface{} {
	var arrayInterface []interface{}
	for i := 0; i < len(Autores); i++ {
		var autorInterface interface{}
		autorInterface = Autores[i]
		arrayInterface = append(arrayInterface, autorInterface)
	}
	return arrayInterface
}
