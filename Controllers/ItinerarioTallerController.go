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

//CreateItinerarioTaller : Metodo de insercion de un nuevo itinerarioTaller
func CreateItinerarioTaller(writter http.ResponseWriter, request *http.Request) {
	var itinerarioTaller models.ItinerarioTaller
	err := json.NewDecoder(request.Body).Decode(&itinerarioTaller)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	itinerarioTallerStrings,itinerarioTallerValues := utilities.ObjectFields(itinerarioTaller)

	result, err := utilities.InsertObject("ItinerarioTaller", itinerarioTallerValues, itinerarioTallerStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "ItinerarioTaller Creada")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetItinerarioTaller : metodo que retorna una relacion Itinerario-Exposicion
func GetItinerarioTaller(writter http.ResponseWriter, request *http.Request) {
	var ItinerarioTaller models.ItinerarioTaller
	err := json.NewDecoder(request.Body).Decode(&ItinerarioTaller)
	jsonResponse := simplejson.New()

	if err == nil {
		var valuesHorario interface{} = nil
		if ItinerarioTaller.Horario != "" {
			valuesHorario = ItinerarioTaller.Horario
		}
		ItinerarioTallerRows, err := utilities.CallStorageProcedure("PAItinerarioTaller", []interface{}{ItinerarioTaller.IDItinerario, ItinerarioTaller.IDTaller, valuesHorario})
		if err == nil {
			var ItinerarioTallerResultado []map[string]interface{}

			if ItinerarioTaller.IDTaller == 0 {

				ItinerarioTallerResultado, err = ItinerariowithTaller(ItinerarioTallerRows)

			} else if ItinerarioTaller.IDItinerario == 0 {

				ItinerarioTallerResultado, err = TallerWithItinerarios(ItinerarioTallerRows)
			}else{

				ItinerarioTallerResultado, err = ItinerariowithTaller(ItinerarioTallerRows)
			}

			if err == nil {
				if len(ItinerarioTallerResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "ItinerarioTaller encontrado")
					jsonResponse.Set("ItinerarioTaller", ItinerarioTallerResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron Itinerario-Taller")
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

//UpdateItinerarioTaller : Metodo que actualiza ItinerarioTaller segun parametros
func UpdateItinerarioTaller(writter http.ResponseWriter, request *http.Request) {
	var lastItinerarioTaller models.ItinerarioTaller
	var newItinerarioTaller models.ItinerarioTaller
	var recipient map[string]interface{}
	err := json.NewDecoder(request.Body).Decode(&recipient)
	jsonResponse := simplejson.New()
	if err == nil {

		mapstructure.Decode(recipient["filter"], &lastItinerarioTaller)
		mapstructure.Decode(recipient["update"], &newItinerarioTaller)

		ItinerarioTallerFilters,ItinerarioTallerFiltersValues :=utilities.ObjectFields(lastItinerarioTaller)

		ItinerarioTallerStrings,ItinerarioTallerValues := utilities.ObjectFields(newItinerarioTaller)

		ItinerarioTallerRows, err := utilities.UpdateObject("ItinerarioTaller", ItinerarioTallerFilters, ItinerarioTallerFiltersValues, ItinerarioTallerStrings, ItinerarioTallerValues)
		if err == nil {

			if ItinerarioTallerRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "ItinerarioTaller actualizado")

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

//ItinerariowithTaller : metodo que retorna una relacion Itinerario-Taller
func ItinerariowithTaller(result *sql.Rows) ([]map[string]interface{}, error) {
	var TallerAux models.Taller
	var ItinerarioAux models.Itinerario
	var tipoAux models.TiposTalleres
	var Itinerarios []models.Itinerario
	var ItinerarioExposcionAux models.ItinerarioTaller
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&ItinerarioAux.ID,
			&ItinerarioAux.Dia,
			&TallerAux.ID,
			&TallerAux.Nombre,
			&TallerAux.Duracion,
			&TallerAux.Enfoque,
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
				"Talleres": []map[string]interface{}{map[string]interface{}{
					"id":          TallerAux.ID,
					"duracion":    TallerAux.Duracion,
					"nombre":      TallerAux.Nombre,
					"enfoque":     TallerAux.Enfoque,
					"descripcion": tipoAux.Descripcion,
					"Horario":     ItinerarioExposcionAux.Horario,
				}},
			}
			response = append(response, newAutorInfo)
		} else {
			var Talleres []map[string]interface{}
			Talleres = response[index]["Talleres"].([]map[string]interface{})
			response[index]["Talleres"] = append(Talleres, map[string]interface{}{
				"id":          TallerAux.ID,
				"duracion":    TallerAux.Duracion,
				"nombre":      TallerAux.Nombre,
				"enfoque":     TallerAux.Enfoque,
				"descripcion": tipoAux.Descripcion,
				"Horario":     ItinerarioExposcionAux.Horario,
			})
		}
	}
	return response, nil
}

//TallerWithItinerarios : metodo que retorna una relacion Itinerario-Taller
func TallerWithItinerarios(result *sql.Rows) ([]map[string]interface{}, error) {
	var TallerAux models.Taller
	var ItinerarioAux models.Itinerario
	var tipoAux models.TiposTalleres
	var Talleres []models.Taller
	var ItinerarioExposcionAux models.ItinerarioTaller
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&ItinerarioAux.ID,
			&ItinerarioAux.Dia,
			&TallerAux.ID,
			&TallerAux.Nombre,
			&TallerAux.Duracion,
			&TallerAux.Enfoque,
			&tipoAux.Descripcion,
			&ItinerarioExposcionAux.Horario)
		if err != nil {
			return nil, err
		}

		index := utilities.Ιndexof(TalleresToInterfaces(Talleres), TallerAux)
		if index == -1 {
			Talleres = append(Talleres, TallerAux)
			newExposicionInfo := map[string]interface{}{
				"id":          TallerAux.ID,
				"duracion":    TallerAux.Duracion,
				"nombre":      TallerAux.Nombre,
				"enfoque":     TallerAux.Enfoque,
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
