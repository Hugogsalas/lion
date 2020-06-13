package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateTaller : Metodo de insercion de un nuevo Taller
func CreateTaller(writter http.ResponseWriter, request *http.Request) {
	var Taller models.Taller
	
	err := json.NewDecoder(request.Body).Decode(&Taller)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	TallerStrings,TallerValues := utilities.ObjectFields(Taller)

	result, err := utilities.InsertObject("Taller", TallerValues, TallerStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result !=0{
		json.Set("Exito", true)
		json.Set("Message", "Taller creado")
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




//GetTaller : Metodo que regresa Talleres segun parametros
func GetTaller(writter http.ResponseWriter, request *http.Request) {
	var Taller models.Taller
	err := json.NewDecoder(request.Body).Decode(&Taller)
	jsonResponse := simplejson.New()
	if err == nil {

		TallerStrings,TallerValues := utilities.ObjectFields(Taller)
		
		var TallerQuery models.GetQuery
		
		TallerQuery.Tables=[]string{"Taller"}
		TallerQuery.Selects=nil
		TallerQuery.Params=[][]string{TallerStrings}
		TallerQuery.Values=[][]interface{}{TallerValues}
		TallerQuery.Conditions=nil

		TallerRows, err := utilities.GetObject(TallerQuery)
		if err == nil {
			TalleresResultado, err := QueryToTaller(TallerRows)
			if err == nil {
				if len(TalleresResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "Talleres encontrados")
					jsonResponse.Set("Talleres", TalleresResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron Talleres")
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

//UpdateTaller : Metodo que actualiza Taller segun parametros
func UpdateTaller(writter http.ResponseWriter, request *http.Request) {
	var Taller models.Taller
	err := json.NewDecoder(request.Body).Decode(&Taller)
	jsonResponse := simplejson.New()
	if err == nil {

		var TallerFilters []string
		var TallerFiltersValues []interface{}

		TallerFilters = append(TallerFilters, "ID")
		TallerFiltersValues = append(TallerFiltersValues, Taller.ID)

		Taller.ID=0
		
		TallerStrings,TallerValues := utilities.ObjectFields(Taller)

		TallerRows, err := utilities.UpdateObject("Taller", TallerFilters, TallerFiltersValues, TallerStrings, TallerValues)
		if err == nil {

			if TallerRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "Taller actualizado")

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

//DeleteTaller : Metodo que elimina Talleres segun parametros
func DeleteTaller(writter http.ResponseWriter, request *http.Request) {
	var Taller models.Taller
	err := json.NewDecoder(request.Body).Decode(&Taller)
	jsonResponse := simplejson.New()
	if err == nil {

		TallerStrings, TallerValues := utilities.ObjectFields(Taller)

		TallerDel, err := utilities.DeleteObject("Taller", TallerStrings, TallerValues)
		if err == nil {

			if TallerDel {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "Taller eliminado")

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

//QueryToTaller : Metodo que transforma la consulta a objetos Taller
func QueryToTaller(result *sql.Rows) ([]models.Taller, error) {
	var TallerAux models.Taller
	var recipents []models.Taller
	for result.Next() {
		err := result.Scan(
			&TallerAux.ID,
			 &TallerAux.Nombre, 
			 &TallerAux.Enfoque,
			 &TallerAux.Duracion, 
			 &TallerAux.IDTipo)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, TallerAux)
	}
	return recipents, nil
}

//TalleresToInterfaces : metodo que transforma un arreglo de Talleres en interfaces
func TalleresToInterfaces(Talleres []models.Taller) []interface{} {
	var arrayInterface []interface{}
	for i:=0;i<len(Talleres);i++{
		var TallerInterface interface{}
		TallerInterface=Talleres[i]
		arrayInterface=append(arrayInterface,TallerInterface)
	}
	return arrayInterface
}