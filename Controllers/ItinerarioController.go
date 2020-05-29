package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateItinerario : Metodo de insercion de un itinerario
func CreateItinerario(writter http.ResponseWriter, request *http.Request) {
	var itinerario models.Itinerario
	err := json.NewDecoder(request.Body).Decode(&itinerario)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var itinerarioValues []interface{}
	var itinerarioStrings []string
	itinerarioValues = utilities.ObjectValues(itinerario)
	itinerarioStrings = utilities.ObjectFields(itinerario)

	result, err := utilities.InsertObject("itinerario", itinerarioValues, itinerarioStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Itinerario Creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetItinerario : Metodo que regresa Itinerarios segun parametros
func GetItinerario(writter http.ResponseWriter, request *http.Request) {
	var itinerario models.Itinerario
	err := json.NewDecoder(request.Body).Decode(&itinerario)
	jsonResponse := simplejson.New()
	if err == nil {

		var itinValues []interface{}
		var itinStrings []string
		itinValues = utilities.ObjectValues(itinerario)
		itinStrings = utilities.ObjectFields(itinerario)

		fmt.Println(itinStrings)
		fmt.Println(itinValues)

		//Limpia de los atributos del objeto
		
		if itinValues[0] == 0{
			itinValues[0] = nil
		}
		
		if itinValues[1] == "" {
			itinValues[1] = nil
		}
		
	

		itinRows, err := utilities.GetObject([]string{"itinerario"}, nil, itinStrings, itinValues)
		if err == nil {
			itinerariosResultado, err := QueryToItinerario(itinRows)
			fmt.Println(itinerariosResultado)
			if err == nil {
				if len(itinerariosResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "Itinerarios encontrados")
					jsonResponse.Set("Libros", itinerariosResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron itinerarios")
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

//QueryToItinerario : Metodo que transforma la consulta a objetos Itinerario
func QueryToItinerario(result *sql.Rows) ([]models.Itinerario, error) {
	var itinerarioAux models.Itinerario
	var recipents []models.Itinerario
	for result.Next() {
		err := result.Scan(&itinerarioAux.ID, &itinerarioAux.Dia)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, itinerarioAux)
	}
	return recipents, nil
}