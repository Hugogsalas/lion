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

//GetAutor : Metodo que regresa autores segun parametros
func GetAutor(writter http.ResponseWriter, request *http.Request) {
	var autor models.Autor
	err := json.NewDecoder(request.Body).Decode(&autor)
	jsonResponse := simplejson.New()
	if err == nil {

		var autorValues []interface{}
		var autorStrings []string
		autorValues = utilities.ObjectValues(autor)
		autorStrings = utilities.ObjectFields(autor)


		//Limpia de los atributos del objeto
		if autorValues[0] == 0 {
			autorValues[0] = nil
		}

		for i := 1; i < len(autorStrings); i++ {
			if autorValues[i] == "" {
				autorValues[i] = nil
			}
		}

		autorRows, err := utilities.GetObject([]string{"Autor"}, nil, autorStrings, autorValues)
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
