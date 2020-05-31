package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateTipoExposicion : Metodo de insercion de un nuevo Tipo Exposiciones
func CreateTipoExposicion(writter http.ResponseWriter, request *http.Request) {
	var TipoExposicion models.TiposExposicion
	
	err := json.NewDecoder(request.Body).Decode(&TipoExposicion)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var TipoExposicionValues []interface{}
	var TipoExposicionStrings []string
	TipoExposicionValues = utilities.ObjectValues(TipoExposicion)
	TipoExposicionStrings = utilities.ObjectFields(TipoExposicion)

	result, err := utilities.InsertObject("TiposExposicion", TipoExposicionValues, TipoExposicionStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "TipoExposicion creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}


//GetTiposExposicion : Metodo que regresa TiposExposiciones segun parametros
func GetTiposExposicion(writter http.ResponseWriter, request *http.Request) {
	var TiposExposicion models.TiposExposicion
	err := json.NewDecoder(request.Body).Decode(&TiposExposicion)
	jsonResponse := simplejson.New()
	if err == nil {

		var TiposExposicionValues []interface{}
		var TiposExposicionStrings []string
		TiposExposicionValues = utilities.ObjectValues(TiposExposicion)
		TiposExposicionStrings = utilities.ObjectFields(TiposExposicion)


		//Limpia de los atributos del objeto
		if TiposExposicionValues[0] == 0 {
			TiposExposicionValues[0] = nil
		}

		for i := 1; i < len(TiposExposicionStrings); i++ {
			if TiposExposicionValues[i] == "" {
				TiposExposicionValues[i] = nil
			}
		}

		TiposExposicionRows, err := utilities.GetObject([]string{"TiposExposicion"}, nil, TiposExposicionStrings, TiposExposicionValues)
		if err == nil {
			TiposExposicionesResultado, err := QueryToTiposExposicion(TiposExposicionRows)
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

//QueryToTiposExposicion : Metodo que transforma la consulta a objetos TiposExposicion
func QueryToTiposExposicion(result *sql.Rows) ([]models.TiposExposicion, error) {
	var TiposExposicionAux models.TiposExposicion
	var recipents []models.TiposExposicion
	for result.Next() {
		err := result.Scan(&TiposExposicionAux.ID, &TiposExposicionAux.Descripcion)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, TiposExposicionAux)
	}
	return recipents, nil
}

//TiposExposicionesToInterfaces : metodo que transforma un arreglo de TiposExposiciones en interfaces
func TiposExposicionesToInterfaces(TiposExposiciones []models.TiposExposicion) []interface{} {
	var arrayInterface []interface{}
	for i:=0;i<len(TiposExposiciones);i++{
		var TiposExposicionInterface interface{}
		TiposExposicionInterface=TiposExposiciones[i]
		arrayInterface=append(arrayInterface,TiposExposicionInterface)
	}
	return arrayInterface
}