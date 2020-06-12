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

//CreateExposicion : Metodo de insercion de una nueva Exposicion
func CreateExposicion(writter http.ResponseWriter, request *http.Request) {
	var Exposicion models.Exposicion

	err := json.NewDecoder(request.Body).Decode(&Exposicion)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	ExposicionStrings, ExposicionValues := utilities.ObjectFields(Exposicion)

	result, err := utilities.InsertObject("Exposicion", ExposicionValues, ExposicionStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Exposicion creada")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetExposicion : Metodo que regresa exposiciones segun parametros
func GetExposicion(writter http.ResponseWriter, request *http.Request) {
	var exposicion models.Exposicion
	err := json.NewDecoder(request.Body).Decode(&exposicion)
	jsonResponse := simplejson.New()
	if err == nil {

		expStrings, expValues := utilities.ObjectFields(exposicion)

		var exposicionQuery models.GetQuery

		exposicionQuery.Tables = []string{"Exposicion"}
		exposicionQuery.Selects = nil
		exposicionQuery.Params = [][]string{expStrings}
		exposicionQuery.Values = [][]interface{}{expValues}
		exposicionQuery.Conditions = nil

		expRows, err := utilities.GetObject(exposicionQuery)
		if err == nil {
			exposicionesResultado, err := QueryToExposicion(expRows)
			fmt.Println(exposicionesResultado)
			if err == nil {
				if len(exposicionesResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "exposiciones encontradas")
					jsonResponse.Set("Libros", exposicionesResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron exposiciones")
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

//GetExposicionWithType : Metodo que regresa Stanes con el desgloce de la editorial segun parametros
func GetExposicionWithType(writter http.ResponseWriter, request *http.Request) {
	var exposicion models.Exposicion
	err := json.NewDecoder(request.Body).Decode(&exposicion)
	jsonResponse := simplejson.New()
	if err == nil {

		expStrings, expValues := utilities.ObjectFields(exposicion)

		var exposicionQuery models.GetQuery

		exposicionQuery.Tables = []string{"Exposicion", "TiposExposicion"}
		exposicionQuery.Selects = [][]string{[]string{"ID", "Duracion", "Titulo", "Presentador"}, []string{"Descripcion"}}
		exposicionQuery.Params = [][]string{expStrings}
		exposicionQuery.Values = [][]interface{}{expValues}
		exposicionQuery.Conditions = []string{"TiposExposicion.ID=Exposicion.IDTipo"}

		expRows, err := utilities.GetObject(exposicionQuery)
		if err == nil {
			exposicionesResultado, err := QueryToFullExposicion(expRows)
			fmt.Println(exposicionesResultado)
			if err == nil {
				if len(exposicionesResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "exposiciones encontradas")
					jsonResponse.Set("Exposiciones", exposicionesResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron exposiciones")
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

//UpdateExposicion : Metodo que actualiza exposicion segun parametros
func UpdateExposicion(writter http.ResponseWriter, request *http.Request) {
	var Exposicion models.Exposicion
	err := json.NewDecoder(request.Body).Decode(&Exposicion)
	jsonResponse := simplejson.New()
	if err == nil {

		var ExposicionFilters []string
		var ExposicionFiltersValues []interface{}

		ExposicionFilters = append(ExposicionFilters, "ID")
		ExposicionFiltersValues = append(ExposicionFiltersValues, Exposicion.ID)

		Exposicion.ID=0
		ExposicionStrings,ExposicionValues := utilities.ObjectFields(Exposicion)

		ExposicionRows, err := utilities.UpdateObject("Exposicion", ExposicionFilters, ExposicionFiltersValues, ExposicionStrings, ExposicionValues)
		if err == nil {

			if ExposicionRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "Exposicion actualizada")

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

//QueryToExposicion : Metodo que transforma la consulta a objetos Exposicion
func QueryToExposicion(result *sql.Rows) ([]models.Exposicion, error) {
	var exposicionAux models.Exposicion
	var recipents []models.Exposicion
	for result.Next() {
		err := result.Scan(&exposicionAux.ID, &exposicionAux.Titulo, &exposicionAux.Presentador, &exposicionAux.Duracion, &exposicionAux.IDTipo)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, exposicionAux)
	}
	return recipents, nil
}

//QueryToFullExposicion : Metodo que transforma la consulta a objetos Exposicion
func QueryToFullExposicion(result *sql.Rows) ([]map[string]interface{}, error) {
	var exposicionAux models.Exposicion
	var tipoExpAux models.TiposExposicion
	var recipents []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&exposicionAux.ID,
			&exposicionAux.Duracion,
			&exposicionAux.Titulo,
			&exposicionAux.Presentador,
			&tipoExpAux.Descripcion)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, map[string]interface{}{
			"id":          exposicionAux.ID,
			"duracion":    exposicionAux.Duracion,
			"titulo":      exposicionAux.Titulo,
			"presentador": exposicionAux.Presentador,
			"descripcion": tipoExpAux.Descripcion,
		})
	}
	return recipents, nil
}

//ExposicionToInterfaces : metodo que transforma un arreglo de Exposicion en interfaces
func ExposicionToInterfaces(Exposicion []models.Exposicion) []interface{} {
	var arrayInterface []interface{}
	for i := 0; i < len(Exposicion); i++ {
		var ExposicionInterface interface{}
		ExposicionInterface = Exposicion[i]
		arrayInterface = append(arrayInterface, ExposicionInterface)
	}
	return arrayInterface
}
