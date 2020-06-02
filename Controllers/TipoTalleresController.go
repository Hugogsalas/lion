package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateTiposTalleres : Metodo de insercion de un nuevo Tipo Taller
func CreateTiposTalleres(writter http.ResponseWriter, request *http.Request) {
	var TiposTalleres models.TiposTalleres
	
	err := json.NewDecoder(request.Body).Decode(&TiposTalleres)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var TiposTalleresValues []interface{}
	var TiposTalleresStrings []string
	TiposTalleresValues = utilities.ObjectValues(TiposTalleres)
	TiposTalleresStrings = utilities.ObjectFields(TiposTalleres)

	result, err := utilities.InsertObject("TiposTalleres", TiposTalleresValues, TiposTalleresStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "TiposTalleres creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}
//GetTiposTalleres : Metodo que regresa TiposTallereses segun parametros
func GetTiposTalleres(writter http.ResponseWriter, request *http.Request) {
	var TiposTalleres models.TiposTalleres
	err := json.NewDecoder(request.Body).Decode(&TiposTalleres)
	jsonResponse := simplejson.New()
	if err == nil {

		var TiposTalleresValues []interface{}
		var TiposTalleresStrings []string
		TiposTalleresValues = utilities.ObjectValues(TiposTalleres)
		TiposTalleresStrings = utilities.ObjectFields(TiposTalleres)


		//Limpia de los atributos del objeto
		if TiposTalleresValues[0] == 0 {
			TiposTalleresValues[0] = nil
		}

		for i := 1; i < len(TiposTalleresStrings); i++ {
			if TiposTalleresValues[i] == "" {
				TiposTalleresValues[i] = nil
			}
		}

		TiposTalleresRows, err := utilities.GetObject([]string{"TiposTalleres"}, nil, TiposTalleresStrings, TiposTalleresValues)
		if err == nil {
			TiposTalleresesResultado, err := QueryToTiposTalleres(TiposTalleresRows)
			if err == nil {
				if len(TiposTalleresesResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "TiposTallereses encontrados")
					jsonResponse.Set("TiposTallereses", TiposTalleresesResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron TiposTallereses")
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

//QueryToTiposTalleres : Metodo que transforma la consulta a objetos TiposTalleres
func QueryToTiposTalleres(result *sql.Rows) ([]models.TiposTalleres, error) {
	var TiposTalleresAux models.TiposTalleres
	var recipents []models.TiposTalleres
	for result.Next() {
		err := result.Scan(&TiposTalleresAux.ID, &TiposTalleresAux.Descripcion )
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, TiposTalleresAux)
	}
	return recipents, nil
}

//TiposTalleresesToInterfaces : metodo que transforma un arreglo de TiposTallereses en interfaces
func TiposTalleresesToInterfaces(TiposTallereses []models.TiposTalleres) []interface{} {
	var arrayInterface []interface{}
	for i:=0;i<len(TiposTallereses);i++{
		var TiposTalleresInterface interface{}
		TiposTalleresInterface=TiposTallereses[i]
		arrayInterface=append(arrayInterface,TiposTalleresInterface)
	}
	return arrayInterface
}