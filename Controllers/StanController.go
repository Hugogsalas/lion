package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateStan : Metodo de insercion de un nuevo Stan
func CreateStan(writter http.ResponseWriter, request *http.Request) {
	var Stan models.Stan
	
	err := json.NewDecoder(request.Body).Decode(&Stan)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var StanValues []interface{}
	var StanStrings []string
	StanValues = utilities.ObjectValues(Stan)
	StanStrings = utilities.ObjectFields(Stan)

	result, err := utilities.InsertObject("Stan", StanValues, StanStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Stan creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}


//GetStan : Metodo que regresa Stanes segun parametros
func GetStan(writter http.ResponseWriter, request *http.Request) {
	var Stan models.Stan
	err := json.NewDecoder(request.Body).Decode(&Stan)
	jsonResponse := simplejson.New()
	if err == nil {

		var StanValues []interface{}
		var StanStrings []string
		StanValues = utilities.ObjectValues(Stan)
		StanStrings = utilities.ObjectFields(Stan)


		//Limpia de los atributos del objeto
	
		for i := 0; i < 3; i++ {
			if StanValues[i] == 0 {
				StanValues[i] = nil
			}
		}

		StanRows, err := utilities.GetObject([]string{"Stan"}, nil, StanStrings, StanValues)
		if err == nil {
			StanesResultado, err := QueryToStan(StanRows)
			if err == nil {
				if len(StanesResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "Stanes encontrados")
					jsonResponse.Set("Stanes", StanesResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron Stanes")
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

//QueryToStan : Metodo que transforma la consulta a objetos Stan
func QueryToStan(result *sql.Rows) ([]models.Stan, error) {
	var StanAux models.Stan
	var recipents []models.Stan
	for result.Next() {
		err := result.Scan(&StanAux.ID, &StanAux.IDEditorial, &StanAux.Numero )
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, StanAux)
	}
	return recipents, nil
}

//StanesToInterfaces : metodo que transforma un arreglo de Stanes en interfaces
func StanesToInterfaces(Stanes []models.Stan) []interface{} {
	var arrayInterface []interface{}
	for i:=0;i<len(Stanes);i++{
		var StanInterface interface{}
		StanInterface=Stanes[i]
		arrayInterface=append(arrayInterface,StanInterface)
	}
	return arrayInterface
}