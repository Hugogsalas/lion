package controllers

import (
	"database/sql"
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)



//CreateExposicion : Metodo de insercion de una nueva Exposicion
func CreateExposicion(writter http.ResponseWriter, request *http.Request) {
	var Exposicion models.Exposicion
	
	err := json.NewDecoder(request.Body).Decode(&Exposicion)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var ExposicionValues []interface{}
	var ExposicionStrings []string
	ExposicionValues = utilities.ObjectValues(Exposicion)
	ExposicionStrings = utilities.ObjectFields(Exposicion)

	result, err := utilities.InsertObject("Exposicion", ExposicionValues, ExposicionStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Exposicion creada")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetExposicion : Metodo que regresa exposiciones segun parametros
func GetExposicion(writter http.ResponseWriter, request *http.Request) {
	var exposicion models.Exposicion
	err := json.NewDecoder(request.Body).Decode(&exposicion)
	jsonResponse := simplejson.New()
	if err == nil {

		var expValues []interface{}
		var expStrings []string
		expValues = utilities.ObjectValues(exposicion)
		expStrings = utilities.ObjectFields(exposicion)

		fmt.Println(expStrings)
		fmt.Println(expValues)

		//Limpia de los atributos del objeto
		for i := 0; i < 3; i++ {
			if expValues[i] == 0{
				expValues[i] = nil
			}
		}

		for i := 3; i < 5; i++ {
			if expValues[i] == "" {
				expValues[i] = nil
			}
		}
	


		expRows, err := utilities.GetObject([]string{"Exposicion"}, nil, expStrings, expValues)
		if err == nil {
			exposicionesResultado, err := QueryToExposicion(expRows)
			fmt.Println(exposicionesResultado)
			if err == nil {
				if len(exposicionesResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "exposiciones encontradas")
					jsonResponse.Set("Libros", exposicionesResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron exposiciones")
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

//QueryToExposicion : Metodo que transforma la consulta a objetos Exposicion
func QueryToExposicion(result *sql.Rows) ([]models.Exposicion, error) {
	var exposicionAux models.Exposicion
	var recipents []models.Exposicion
	for result.Next() {
		err := result.Scan(&exposicionAux.ID, &exposicionAux.Titulo, &exposicionAux.Presentador, &exposicionAux.Duracion, &exposicionAux.IDTipo)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, exposicionAux)
	}
	return recipents, nil
}