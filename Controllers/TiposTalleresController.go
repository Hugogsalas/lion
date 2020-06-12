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

	TiposTalleresStrings,TiposTalleresValues := utilities.ObjectFields(TiposTalleres)

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

		TiposTalleresStrings,TiposTalleresValues := utilities.ObjectFields(TiposTalleres)

		var TiposTalleresQuery models.GetQuery
		
		TiposTalleresQuery.Tables=[]string{"TiposTalleres"}
		TiposTalleresQuery.Selects=nil
		TiposTalleresQuery.Params=[][]string{TiposTalleresStrings}
		TiposTalleresQuery.Values=[][]interface{}{TiposTalleresValues}
		TiposTalleresQuery.Conditions=nil

		TiposTalleresRows, err := utilities.GetObject(TiposTalleresQuery)
		if err == nil {
			TiposTalleresesResultado, err := QueryToTiposTalleres(TiposTalleresRows)
			if err == nil {
				if len(TiposTalleresesResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "TiposTallereses encontrados")
					jsonResponse.Set("TiposTalleres", TiposTalleresesResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron TiposTalleres")
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

//UpdateTiposTaller : Metodo que actualiza TiposTaller segun parametros
func UpdateTiposTaller(writter http.ResponseWriter, request *http.Request) {
	var TiposTaller models.TiposTalleres
	err := json.NewDecoder(request.Body).Decode(&TiposTaller)
	jsonResponse := simplejson.New()
	if err == nil {

		var TiposTallerFilters []string
		var TiposTallerFiltersValues []interface{}

		TiposTallerFilters = append(TiposTallerFilters, "ID")
		TiposTallerFiltersValues = append(TiposTallerFiltersValues, TiposTaller.ID)

		TiposTaller.ID=0

		TiposTallerStrings,TiposTallerValues := utilities.ObjectFields(TiposTaller)

		TiposTallerRows, err := utilities.UpdateObject("TiposTalleres", TiposTallerFilters, TiposTallerFiltersValues, TiposTallerStrings, TiposTallerValues)
		if err == nil {

			if TiposTallerRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "TiposTaller actualizado")

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

//DeleteTiposTalleres : Metodo que elimina TiposTalleres segun parametros
func DeleteTiposTalleres(writter http.ResponseWriter, request *http.Request) {
	var TiposTaller models.TiposTalleres
	err := json.NewDecoder(request.Body).Decode(&TiposTaller)
	jsonResponse := simplejson.New()
	if err == nil {

		TiposTallerStrings, TiposTallerValues := utilities.ObjectFields(TiposTaller)

		TiposTallerDel, err := utilities.DeleteObject("TiposTaller", TiposTallerStrings, TiposTallerValues)
		if err == nil {

			if TiposTallerDel {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "TiposTaller eliminado")

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