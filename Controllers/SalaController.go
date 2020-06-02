package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateSala : Metodo de insercion de una nueva sala
func CreateSala(writter http.ResponseWriter, request *http.Request) {
	var sala models.Sala
	err := json.NewDecoder(request.Body).Decode(&sala)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var salaValues []interface{}
	var salaStrings []string
	salaValues = utilities.ObjectValues(sala)
	salaStrings = utilities.ObjectFields(sala)

	result, err := utilities.InsertObject("sala", salaValues, salaStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Sala Creada")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}


//GetSala : Metodo que regresa Salas segun parametros
func GetSala(writter http.ResponseWriter, request *http.Request) {
	var Sala models.Sala
	err := json.NewDecoder(request.Body).Decode(&Sala)
	jsonResponse := simplejson.New()
	if err == nil {

		var SalaValues []interface{}
		var SalaStrings []string
		SalaValues = utilities.ObjectValues(Sala)
		SalaStrings = utilities.ObjectFields(Sala)


		//Limpia de los atributos del objeto
	
		if SalaValues[0] == 0 {
			SalaValues[0] = nil
		}
		if SalaValues[1] == "" {
			SalaValues[1] = nil
		}

		SalaRows, err := utilities.GetObject([]string{"Sala"}, nil, SalaStrings, SalaValues)
		if err == nil {
			SalasResultado, err := QueryToSala(SalaRows)
			if err == nil {
				if len(SalasResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "Salas encontrados")
					jsonResponse.Set("Salas", SalasResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron Salas")
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

//QueryToSala : Metodo que transforma la consulta a objetos Sala
func QueryToSala(result *sql.Rows) ([]models.Sala, error) {
	var SalaAux models.Sala
	var recipents []models.Sala
	for result.Next() {
		err := result.Scan(&SalaAux.ID, &SalaAux.Nombre )
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, SalaAux)
	}
	return recipents, nil
}

//SalasToInterfaces : metodo que transforma un arreglo de Salas en interfaces
func SalasToInterfaces(Salas []models.Sala) []interface{} {
	var arrayInterface []interface{}
	for i:=0;i<len(Salas);i++{
		var SalaInterface interface{}
		SalaInterface=Salas[i]
		arrayInterface=append(arrayInterface,SalaInterface)
	}
	return arrayInterface
}
