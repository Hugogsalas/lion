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

//CreateSalaExposicion : Metodo de insercion de una nueva relacion Sala-Exposicion
func CreateSalaExposicion(writter http.ResponseWriter, request *http.Request) {
	var SalaExposicion models.SalaExposicion

	err := json.NewDecoder(request.Body).Decode(&SalaExposicion)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	SalaExposicionestrings, SalaExposicionValues := utilities.ObjectFields(SalaExposicion)

	result, err := utilities.InsertObject("SalaExposicion", SalaExposicionValues, SalaExposicionestrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result==0 && err==nil {
		json.Set("Exito", true)
		json.Set("Message", "Sala-Exposicion creado")
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

//GetSalaExposicion : metodo que retorna  una relacion Sala-Exposicion
func GetSalaExposicion(writter http.ResponseWriter, request *http.Request) {
	var SalaExposicion models.SalaExposicion
	err := json.NewDecoder(request.Body).Decode(&SalaExposicion)
	jsonResponse := simplejson.New()

	if err == nil {

		SalaExposicionRows, err := utilities.CallStorageProcedure("PASalaExposicion", []interface{}{SalaExposicion.IDSala, SalaExposicion.IDExposicion})
		if err == nil {
			var SalaExposicionResultado []map[string]interface{}

			if SalaExposicion.IDExposicion == 0 {
				SalaExposicionResultado, err = SalaWithExposiciones(SalaExposicionRows)
			} else {
				SalaExposicionResultado, err = ExposicionesWithSala(SalaExposicionRows)
			}

			if err == nil {
				if len(SalaExposicionResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "ExposicionExposicion encontrado")
					jsonResponse.Set("SalaExposicion", SalaExposicionResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron SalaExposiciones")
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

//UpdateSalaExposicion : Metodo que actualiza SalaExposicion segun parametros
func UpdateSalaExposicion(writter http.ResponseWriter, request *http.Request) {
	var lastSalaExposicion models.SalaExposicion
	var newSalaExposicion models.SalaExposicion
	var recipient map[string]interface{}
	err := json.NewDecoder(request.Body).Decode(&recipient)
	jsonResponse := simplejson.New()
	if err == nil {

		mapstructure.Decode(recipient["filter"], &lastSalaExposicion)
		mapstructure.Decode(recipient["update"], &newSalaExposicion)

		SalaExposicionFilters, SalaExposicionFiltersValues := utilities.ObjectFields(lastSalaExposicion)
		SalaExposicionStrings, SalaExposicionValues := utilities.ObjectFields(newSalaExposicion)

		SalaExposicionRows, err := utilities.UpdateObject("SalaExposicion", SalaExposicionFilters, SalaExposicionFiltersValues, SalaExposicionStrings, SalaExposicionValues)
		if err == nil {

			if SalaExposicionRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "SalaExposicion actualizado")

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

//DeleteSalaExposicion : Metodo que elimina SalaExposiciones segun parametros
func DeleteSalaExposicion(writter http.ResponseWriter, request *http.Request) {
	var SalaExposicion models.SalaExposicion
	err := json.NewDecoder(request.Body).Decode(&SalaExposicion)
	jsonResponse := simplejson.New()
	if err == nil {

		SalaExposicionStrings, SalaExposicionValues := utilities.ObjectFields(SalaExposicion)

		SalaExposicionDel, err := utilities.DeleteObject("SalaExposicion", SalaExposicionStrings, SalaExposicionValues)
		if err == nil {

			if SalaExposicionDel {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "SalaExposicion eliminado")

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

//SalaWithExposiciones : metodo que combierte una consulta a una relacion Sala con Exposiciones descritos
func SalaWithExposiciones(result *sql.Rows) ([]map[string]interface{}, error) {
	var ExposicionAux models.Exposicion
	var tipoAux models.TiposExposiciones
	var SalaAux models.Sala
	var Salas []models.Sala
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&SalaAux.ID,
			&SalaAux.Nombre,
			&ExposicionAux.ID,
			&ExposicionAux.Presentador,
			&ExposicionAux.Titulo,
			&ExposicionAux.Duracion,
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
				"Exposiciones": []map[string]interface{}{map[string]interface{}{
					"id":          ExposicionAux.ID,
					"duracion":    ExposicionAux.Duracion,
					"descripcion": tipoAux.Descripcion,
					"titulo":      ExposicionAux.Titulo,
					"presentador": ExposicionAux.Presentador,
				}},
			}
			response = append(response, newSalaInfo)
		} else {
			var lastExposiciones []map[string]interface{}
			lastExposiciones = response[index]["Exposiciones"].([]map[string]interface{})
			response[index]["Exposiciones"] = append(lastExposiciones, map[string]interface{}{
				"id":          ExposicionAux.ID,
				"duracion":    ExposicionAux.Duracion,
				"descripcion": tipoAux.Descripcion,
				"titulo":      ExposicionAux.Titulo,
				"presentador": ExposicionAux.Presentador,
			})
		}
	}
	return response, nil
}

//ExposicionesWithSala : metodo que combierte una consulta a una relacion libors con Salas descritos
func ExposicionesWithSala(result *sql.Rows) ([]map[string]interface{}, error) {
	var ExposicionAux models.Exposicion
	var tipoAux models.TiposExposiciones
	var SalaAux models.Sala
	var Exposiciones []models.Exposicion
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&SalaAux.ID,
			&SalaAux.Nombre,
			&ExposicionAux.ID,
			&ExposicionAux.Presentador,
			&ExposicionAux.Titulo,
			&ExposicionAux.Duracion,
			&tipoAux.Descripcion)
		if err != nil {
			return nil, err
		}

		index := utilities.Ιndexof(ExposicionToInterfaces(Exposiciones), ExposicionAux)
		if index == -1 {
			Exposiciones = append(Exposiciones, ExposicionAux)
			newExposicionInfo := map[string]interface{}{
				"id":          ExposicionAux.ID,
				"Duracion":    ExposicionAux.Duracion,
				"Titulo":      ExposicionAux.Titulo,
				"Presentador": ExposicionAux.Presentador,
				"descripcion": tipoAux.Descripcion,
				"Salas":       []models.Sala{SalaAux},
			}
			response = append(response, newExposicionInfo)
		} else {
			var Salas []models.Sala
			Salas = response[index]["Salas"].([]models.Sala)
			response[index]["Salas"] = append(Salas, SalaAux)
		}
	}
	return response, nil
}
