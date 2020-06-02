package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/mitchellh/mapstructure"

	models "../Models"
	utilities "../Utilities"
)

//CreateItinerarioExposicion : Metodo de insercion de un itinerarioExposicion
func CreateItinerarioExposicion(writter http.ResponseWriter, request *http.Request) {
	var itinerarioExposicion models.ItinerarioExposicion
	err := json.NewDecoder(request.Body).Decode(&itinerarioExposicion)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var itinerarioExposicionValues []interface{}
	var itinerarioExposicionStrings []string
	itinerarioExposicionValues = utilities.ObjectValues(itinerarioExposicion)
	itinerarioExposicionStrings = utilities.ObjectFields(itinerarioExposicion)

	result, err := utilities.InsertObject("itinerarioExposicion", itinerarioExposicionValues, itinerarioExposicionStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "ItinerarioExposicion Creada")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetItinerarioExposicion : metodo que retorna una relacion Itinerario-Exposicion
func GetItinerarioExposicion(writter http.ResponseWriter, request *http.Request) {
	var ItinerarioExposicion models.ItinerarioExposicion
	err := json.NewDecoder(request.Body).Decode(&ItinerarioExposicion)
	jsonResponse := simplejson.New()

	if err == nil {
		var valuesHorario interface{} = nil
		if ItinerarioExposicion.Horario != "" {
			valuesHorario = ItinerarioExposicion.Horario
		}
		ItinerarioExposicionRows, err := utilities.CallStorageProcedure("PAItinerarioExposicion", []interface{}{ItinerarioExposicion.IDItinerario, ItinerarioExposicion.IDExposicion, valuesHorario})
		if err == nil {
			var ItinerarioExposicionResultado []map[string]interface{}

			if ItinerarioExposicion.IDExposicion == 0 {

				ItinerarioExposicionResultado, err = ItinerariowithExposicion(ItinerarioExposicionRows)

			} else if ItinerarioExposicion.IDItinerario == 0 {

				ItinerarioExposicionResultado, err = ExposicionWithItinerarios(ItinerarioExposicionRows)
			} else {
				ItinerarioExposicionResultado, err = ItinerariowithExposicion(ItinerarioExposicionRows)
			}

			if err == nil {
				if len(ItinerarioExposicionResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "ItinerarioExposicion encontrado")
					jsonResponse.Set("ItinerarioExposicion", ItinerarioExposicionResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron Itinerario-Exposicion")
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

//UpdateItinerarioExposicion : Metodo que actualiza ItinerarioExposicion segun parametros
func UpdateItinerarioExposicion(writter http.ResponseWriter, request *http.Request) {
	var lastItinerarioExposicion models.ItinerarioExposicion
	var newItinerarioExposicion models.ItinerarioExposicion
	var recipient map[string]interface{}
	err := json.NewDecoder(request.Body).Decode(&recipient)
	jsonResponse := simplejson.New()
	if err == nil {

		mapstructure.Decode(recipient["filter"], &lastItinerarioExposicion)
		mapstructure.Decode(recipient["update"], &newItinerarioExposicion)

		var ItinerarioExposicionFiltersValues []interface{}
		var ItinerarioExposicionFilters []string

		ItinerarioExposicionFiltersValues =utilities.ObjectValues(lastItinerarioExposicion)
		ItinerarioExposicionFilters =utilities.ObjectFields(lastItinerarioExposicion)

		var ItinerarioExposicionValues []interface{}
		var ItinerarioExposicionStrings []string

		ItinerarioExposicionValues = utilities.ObjectValues(newItinerarioExposicion)
		ItinerarioExposicionStrings = utilities.ObjectFields(newItinerarioExposicion)

		for i:=1;i<2;i++{
			if ItinerarioExposicionValues[i] == 0 {
				ItinerarioExposicionValues[i] = nil
			}
			if ItinerarioExposicionFiltersValues[i] == 0 {
				ItinerarioExposicionFiltersValues[i] = nil
			}
		}
		if ItinerarioExposicionValues[2] == "" {
			ItinerarioExposicionValues[2] = nil
		}
		if ItinerarioExposicionFiltersValues[2] == "" {
			ItinerarioExposicionFiltersValues[2] = nil
		}

		ItinerarioExposicionRows, err := utilities.UpdateObject("ItinerarioExposicion", ItinerarioExposicionFilters, ItinerarioExposicionFiltersValues, ItinerarioExposicionStrings, ItinerarioExposicionValues)
		if err == nil {

			if ItinerarioExposicionRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "ItinerarioExposicion actualizado")

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

//ItinerariowithExposicion : metodo que retorna una relacion Itinerario-Exposicion
func ItinerariowithExposicion(result *sql.Rows) ([]map[string]interface{}, error) {
	var ExposicionAux models.Exposicion
	var ItinerarioAux models.Itinerario
	var tipoAux models.TiposExposicion
	var Itinerarios []models.Itinerario
	var ItinerarioExposcionAux models.ItinerarioExposicion
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&ItinerarioAux.ID,
			&ItinerarioAux.Dia,
			&ExposicionAux.ID,
			&ExposicionAux.Presentador,
			&ExposicionAux.Duracion,
			&ExposicionAux.Titulo,
			&tipoAux.Descripcion,
			&ItinerarioExposcionAux.Horario)
		if err != nil {
			return nil, err
		}

		index := utilities.Ιndexof(ItinerariosToInterfaces(Itinerarios), ItinerarioAux)
		if index == -1 {
			Itinerarios = append(Itinerarios, ItinerarioAux)
			newAutorInfo := map[string]interface{}{
				"id":  ItinerarioAux.ID,
				"dia": ItinerarioAux.Dia,
				"Exposiciones": []map[string]interface{}{map[string]interface{}{
					"id":          ExposicionAux.ID,
					"duracion":    ExposicionAux.Duracion,
					"titulo":      ExposicionAux.Titulo,
					"presentador": ExposicionAux.Presentador,
					"descripcion": tipoAux.Descripcion,
					"Horario":     ItinerarioExposcionAux.Horario,
				}},
			}
			response = append(response, newAutorInfo)
		} else {
			var Exposiciones []map[string]interface{}
			Exposiciones = response[index]["Exposiciones"].([]map[string]interface{})
			response[index]["Exposiciones"] = append(Exposiciones, map[string]interface{}{
				"id":          ExposicionAux.ID,
				"duracion":    ExposicionAux.Duracion,
				"titulo":      ExposicionAux.Titulo,
				"presentador": ExposicionAux.Presentador,
				"descripcion": tipoAux.Descripcion,
				"Horario":     ItinerarioExposcionAux.Horario,
			})
		}
	}
	return response, nil
}

//ExposicionWithItinerarios : metodo que retorna una relacion Itinerario-Exposicion
func ExposicionWithItinerarios(result *sql.Rows) ([]map[string]interface{}, error) {
	var ExposicionAux models.Exposicion
	var ItinerarioAux models.Itinerario
	var tipoAux models.TiposExposicion
	var Exposiciones []models.Exposicion
	var ItinerarioExposcionAux models.ItinerarioExposicion
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&ItinerarioAux.ID,
			&ItinerarioAux.Dia,
			&ExposicionAux.ID,
			&ExposicionAux.Presentador,
			&ExposicionAux.Duracion,
			&ExposicionAux.Titulo,
			&tipoAux.Descripcion,
			&ItinerarioExposcionAux.Horario)
		if err != nil {
			return nil, err
		}

		index := utilities.Ιndexof(ExposicionToInterfaces(Exposiciones), ExposicionAux)
		if index == -1 {
			Exposiciones = append(Exposiciones, ExposicionAux)
			newExposicionInfo := map[string]interface{}{
				"id":          ExposicionAux.ID,
				"duracion":    ExposicionAux.Duracion,
				"titulo":      ExposicionAux.Titulo,
				"presentador": ExposicionAux.Presentador,
				"descripcion": tipoAux.Descripcion,
				"Itinerarios": []map[string]interface{}{map[string]interface{}{
					"id":      ItinerarioAux.ID,
					"dia":     ItinerarioAux.Dia,
					"Horario": ItinerarioExposcionAux.Horario,
				}},
			}
			response = append(response, newExposicionInfo)
		} else {
			var Itinerarios []map[string]interface{}
			Itinerarios = response[index]["Itinerarios"].([]map[string]interface{})
			response[index]["Itinerarios"] = append(Itinerarios, map[string]interface{}{
				"id":      ItinerarioAux.ID,
				"dia":     ItinerarioAux.Dia,
				"Horario": ItinerarioExposcionAux.Horario,
			})
		}
	}
	return response, nil
}
