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

//CreateSalaTaller : Metodo de insercion de una nueva relacion Sala-Taller
func CreateSalaTaller(writter http.ResponseWriter, request *http.Request) {
	var SalaTaller models.SalaTaller
	
	err := json.NewDecoder(request.Body).Decode(&SalaTaller)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}
	SalaTallerStrings,SalaTallerValues := utilities.ObjectFields(SalaTaller)

	result, err := utilities.InsertObject("SalaTaller", SalaTallerValues, SalaTallerStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Sala-Taller creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}


//GetSalaTaller : metodo que retorna  una relacion Sala-Taller
func GetSalaTaller(writter http.ResponseWriter, request *http.Request) {
	var SalaTaller models.SalaTaller
	err := json.NewDecoder(request.Body).Decode(&SalaTaller)
	jsonResponse := simplejson.New()

	if err == nil {

		SalaTallerRows, err := utilities.CallStorageProcedure("PASalaTaller", []interface{}{SalaTaller.IDSala, SalaTaller.IDTaller})
		if err == nil {
			var SalaTallerResultado []map[string]interface{}

			if SalaTaller.IDTaller == 0 {
				SalaTallerResultado, err = SalaWithTalleres(SalaTallerRows)
			} else {
				SalaTallerResultado, err = TalleresWithSala(SalaTallerRows)
			}

			if err == nil {
				if len(SalaTallerResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "TallerTaller encontrado")
					jsonResponse.Set("SalaTaller", SalaTallerResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron SalaTalleres")
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

//UpdateSalaTaller : Metodo que actualiza SalaTaller segun parametros
func UpdateSalaTaller(writter http.ResponseWriter, request *http.Request) {
	var lastSalaTaller models.SalaTaller
	var newSalaTaller models.SalaTaller
	var recipient map[string]interface{}
	err := json.NewDecoder(request.Body).Decode(&recipient)
	jsonResponse := simplejson.New()
	if err == nil {

		mapstructure.Decode(recipient["filter"], &lastSalaTaller)
		mapstructure.Decode(recipient["update"], &newSalaTaller)

		SalaTallerFilters,SalaTallerFiltersValues :=utilities.ObjectFields(lastSalaTaller)
		SalaTallerStrings,SalaTallerValues := utilities.ObjectFields(newSalaTaller)

		SalaTallerRows, err := utilities.UpdateObject("SalaTaller", SalaTallerFilters, SalaTallerFiltersValues, SalaTallerStrings, SalaTallerValues)
		if err == nil {

			if SalaTallerRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "SalaTaller actualizado")

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

//DeleteSalaTaller : Metodo que elimina SalaTalleres segun parametros
func DeleteSalaTaller(writter http.ResponseWriter, request *http.Request) {
	var SalaTaller models.SalaTaller
	err := json.NewDecoder(request.Body).Decode(&SalaTaller)
	jsonResponse := simplejson.New()
	if err == nil {

		SalaTallerStrings, SalaTallerValues := utilities.ObjectFields(SalaTaller)

		SalaTallerDel, err := utilities.DeleteObject("SalaTaller", SalaTallerStrings, SalaTallerValues)
		if err == nil {

			if SalaTallerDel {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "SalaTaller eliminado")

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

//SalaWithTalleres : metodo que combierte una consulta a una relacion Sala con Talleres descritos
func SalaWithTalleres(result *sql.Rows) ([]map[string]interface{}, error) {
	var TallerAux models.Taller
	var tipoAux models.TiposTalleres
	var SalaAux models.Sala
	var Salas []models.Sala
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&SalaAux.ID,
			&SalaAux.Nombre,
			&TallerAux.ID,
			&TallerAux.Nombre,
			&TallerAux.Enfoque,
			&TallerAux.Duracion,
			&tipoAux.Descripcion)

		if err != nil {
			return nil, err
		}

		index := utilities.Ιndexof(SalasToInterfaces(Salas), SalaAux)
		if index == -1 {
			Salas = append(Salas, SalaAux)
			newSalaInfo := map[string]interface{}{
				"id":     SalaAux.ID,
				"nombre": SalaAux.Nombre,
				"Talleres": []map[string]interface{}{map[string]interface{}{
					"id" :          TallerAux.ID,
					"duracion" : TallerAux.Duracion,
					"descripcion": tipoAux.Descripcion,
					"nombre" :      TallerAux.Nombre,
					"Enfoque" : TallerAux.Enfoque,
				}},
			}
			response = append(response, newSalaInfo)
		} else {
			var lastTalleres []map[string]interface{}
			lastTalleres = response[index]["Talleres"].([]map[string]interface{})
			response[index]["Talleres"] = append(lastTalleres,map[string]interface{}{
				"id":          TallerAux.ID,
				"duracion": TallerAux.Duracion,
				"nombre":      TallerAux.Nombre,
				"Enfoque": TallerAux.Enfoque,
				"descripcion": tipoAux.Descripcion,
			})
		}
	}
	return response, nil
}

//TalleresWithSala : metodo que combierte una consulta a una relacion libors con Salas descritos
func TalleresWithSala(result *sql.Rows) ([]map[string]interface{}, error) {
	var TallerAux models.Taller
	var tipoAux models.TiposTalleres
	var SalaAux models.Sala
	var Talleres []models.Taller
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&SalaAux.ID,
			&SalaAux.Nombre,
			&TallerAux.ID,
			&TallerAux.Nombre,
			&TallerAux.Enfoque,
			&TallerAux.Duracion,
			&tipoAux.Descripcion)
		if err != nil {
			return nil, err
		}

		index := utilities.Ιndexof(TalleresToInterfaces(Talleres), TallerAux)
		if index == -1 {
			Talleres = append(Talleres, TallerAux)
			newTallerInfo := map[string]interface{}{
				"id":          TallerAux.ID,
				"duracion":      TallerAux.Duracion,
				"nombre": TallerAux.Nombre,
				"Enfoque": 	TallerAux.Enfoque,
				"descripcion": tipoAux.Descripcion,
				"Salas":       []models.Sala{SalaAux},
			}
			response = append(response, newTallerInfo)
		} else {
			var Salas []models.Sala
			Salas = response[index]["Salas"].([]models.Sala)
			response[index]["Salas"] = append(Salas, SalaAux)
		}
	}
	return response, nil
}
