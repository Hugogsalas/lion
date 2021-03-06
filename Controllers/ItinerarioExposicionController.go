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

	itinerarioExposicionStrings, itinerarioExposicionValues := utilities.ObjectFields(itinerarioExposicion)

	result, err := utilities.InsertObject("itinerarioExposicion", itinerarioExposicionValues, itinerarioExposicionStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result==0 && err==nil {
		json.Set("Exito", true)
		json.Set("Message", "ItinerarioExposicion Creada")
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

		ItinerarioExposicionFilters, ItinerarioExposicionFiltersValues := utilities.ObjectFields(lastItinerarioExposicion)
		ItinerarioExposicionStrings, ItinerarioExposicionValues := utilities.ObjectFields(newItinerarioExposicion)

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

//DeleteItinerarioExposicion : Metodo que elimina ItinerarioExposicion segun parametros
func DeleteItinerarioExposicion(writter http.ResponseWriter, request *http.Request) {
	var ItinerarioExposicion models.ItinerarioExposicion
	err := json.NewDecoder(request.Body).Decode(&ItinerarioExposicion)
	jsonResponse := simplejson.New()
	if err == nil {

		ItinerarioExposicionStrings, ItinerarioExposicionValues := utilities.ObjectFields(ItinerarioExposicion)

		ItinerarioExposicionDel, err := utilities.DeleteObject("ItinerarioExposicion", ItinerarioExposicionStrings, ItinerarioExposicionValues)
		if err == nil {

			if ItinerarioExposicionDel {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "ItinerarioExposicion eliminado")

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
	var tipoAux models.TiposExposiciones
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
	var tipoAux models.TiposExposiciones
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
