package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateTiposExposiciones : Metodo de insercion de un nuevo Tipo Exposiciones
func CreateTiposExposiciones(writter http.ResponseWriter, request *http.Request) {
	var TiposExposiciones models.TiposExposiciones
	
	err := json.NewDecoder(request.Body).Decode(&TiposExposiciones)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	TiposExposicionesStrings,TiposExposicionesValues := utilities.ObjectFields(TiposExposiciones)

	result, err := utilities.InsertObject("TiposExposiciones", TiposExposicionesValues, TiposExposicionesStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "TiposExposiciones creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}


//GetTiposExposiciones : Metodo que regresa TiposExposiciones segun parametros
func GetTiposExposiciones(writter http.ResponseWriter, request *http.Request) {
	var TiposExposiciones models.TiposExposiciones
	err := json.NewDecoder(request.Body).Decode(&TiposExposiciones)
	jsonResponse := simplejson.New()
	if err == nil {

		TiposExposicionesStrings,TiposExposicionesValues := utilities.ObjectFields(TiposExposiciones)

		var TiposExposicionesQuery models.GetQuery
		
		TiposExposicionesQuery.Tables=[]string{"TiposExposiciones"}
		TiposExposicionesQuery.Selects=nil
		TiposExposicionesQuery.Params=[][]string{TiposExposicionesStrings}
		TiposExposicionesQuery.Values=[][]interface{}{TiposExposicionesValues}
		TiposExposicionesQuery.Conditions=nil

		TiposExposicionesRows, err := utilities.GetObject(TiposExposicionesQuery)
		if err == nil {
			TiposExposicionesResultado, err := QueryToTiposExposiciones(TiposExposicionesRows)
			if err == nil {
				if len(TiposExposicionesResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "TiposExposiciones encontrados")
					jsonResponse.Set("TiposExposiciones", TiposExposicionesResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron TiposExposiciones")
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

//UpdateTiposExposiciones : Metodo que actualiza TiposExposiciones segun parametros
func UpdateTiposExposiciones(writter http.ResponseWriter, request *http.Request) {
	var TiposExposiciones models.TiposExposiciones
	err := json.NewDecoder(request.Body).Decode(&TiposExposiciones)
	jsonResponse := simplejson.New()
	if err == nil {

		var TiposExposicionesFilters []string
		var TiposExposicionesFiltersValues []interface{}

		TiposExposicionesFilters = append(TiposExposicionesFilters, "ID")
		TiposExposicionesFiltersValues = append(TiposExposicionesFiltersValues, TiposExposiciones.ID)
		
		TiposExposiciones.ID=0

		TiposExposicionesStrings,TiposExposicionesValues := utilities.ObjectFields(TiposExposiciones)


		TiposExposicionesRows, err := utilities.UpdateObject("TiposExposiciones", TiposExposicionesFilters, TiposExposicionesFiltersValues, TiposExposicionesStrings, TiposExposicionesValues)
		if err == nil {

			if TiposExposicionesRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "TiposExposiciones actualizado")

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

//DeleteTiposExposiciones : Metodo que elimina TiposExposiciones segun parametros
func DeleteTiposExposiciones(writter http.ResponseWriter, request *http.Request) {
	var TiposExposiciones models.TiposExposiciones
	err := json.NewDecoder(request.Body).Decode(&TiposExposiciones)
	jsonResponse := simplejson.New()
	if err == nil {

		TiposExposicionesStrings, TiposExposicionesValues := utilities.ObjectFields(TiposExposiciones)

		TiposExposicionesDel, err := utilities.DeleteObject("TiposExposiciones", TiposExposicionesStrings, TiposExposicionesValues)
		if err == nil {

			if TiposExposicionesDel {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "TiposExposiciones eliminado")

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


//QueryToTiposExposiciones : Metodo que transforma la consulta a objetos TiposExposiciones
func QueryToTiposExposiciones(result *sql.Rows) ([]models.TiposExposiciones, error) {
	var TiposExposicionesAux models.TiposExposiciones
	var recipents []models.TiposExposiciones
	for result.Next() {
		err := result.Scan(&TiposExposicionesAux.ID, &TiposExposicionesAux.Descripcion)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, TiposExposicionesAux)
	}
	return recipents, nil
}

//TiposExposicionesToInterfaces : metodo que transforma un arreglo de TiposExposiciones en interfaces
func TiposExposicionesToInterfaces(TiposExposiciones []models.TiposExposiciones) []interface{} {
	var arrayInterface []interface{}
	for i:=0;i<len(TiposExposiciones);i++{
		var TiposExposicionesInterface interface{}
		TiposExposicionesInterface=TiposExposiciones[i]
		arrayInterface=append(arrayInterface,TiposExposicionesInterface)
	}
	return arrayInterface
}