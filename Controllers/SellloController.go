package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	
	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateSello : Metodo de insercion de un nuevo sello
func CreateSello(writter http.ResponseWriter, request *http.Request) {
	var Sello models.Sello
	
	err := json.NewDecoder(request.Body).Decode(&Sello)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	SelloStrings,SelloValues := utilities.ObjectFields(Sello)

	result, err := utilities.InsertObject("Sello", SelloValues, SelloStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Sello creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}


//GetSello : Metodo que regresa sello segun parametros
func GetSello(writter http.ResponseWriter, request *http.Request) {
	var Sello models.Sello
	err := json.NewDecoder(request.Body).Decode(&Sello)
	jsonResponse := simplejson.New()

	if err == nil {

		selloStrings,selloValues := utilities.ObjectFields(Sello)

		var SelloQuery models.GetQuery
		
		SelloQuery.Tables=[]string{"Sello"}
		SelloQuery.Selects=nil
		SelloQuery.Params=[][]string{selloStrings}
		SelloQuery.Values=[][]interface{}{selloValues}
		SelloQuery.Conditions=nil

		selloRows, err := utilities.GetObject(SelloQuery)
		if err == nil {
			selloResultado, err := QueryToSello(selloRows)
			if err == nil {
				if len(selloResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "El sello ha sido encontrada")
					jsonResponse.Set("Editorial", selloResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron sellos")
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

//UpdateSello : Metodo que actualiza Sellos segun parametros
func UpdateSello(writter http.ResponseWriter, request *http.Request) {
	var Sello models.Sello
	err := json.NewDecoder(request.Body).Decode(&Sello)
	jsonResponse := simplejson.New()
	if err == nil {

		var SelloFilters []string
		var SelloFiltersValues []interface{}

		SelloFilters = append(SelloFilters, "ID")
		SelloFiltersValues = append(SelloFiltersValues, Sello.ID)

		Sello.ID=0

		SelloStrings,SelloValues := utilities.ObjectFields(Sello)

		SelloRows, err := utilities.UpdateObject("Sello", SelloFilters, SelloFiltersValues, SelloStrings, SelloValues)
		if err == nil {

			if SelloRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "Sello actualizado")

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

//QueryToSello : Metodo que transforma la consulta a objetos Sello
func QueryToSello(result *sql.Rows) ([]models.Sello, error) {
	var selloAux models.Sello
	var recipents []models.Sello
	for result.Next() {
		err := result.Scan(
			&selloAux.ID,
			&selloAux.IDEditorial, 
			&selloAux.Descripcion)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, selloAux)
	}
	return recipents, nil
}

//SellosToInterfaces : metodo que transforma un arreglo de Sellos en interfaces
func SellosToInterfaces(Sellos []models.Sello) []interface{} {
	var arrayInterface []interface{}
	for i:=0;i<len(Sellos);i++{
		var selloInterface interface{}
		selloInterface=Sellos[i]
		arrayInterface=append(arrayInterface,selloInterface)
	}
	return arrayInterface
}