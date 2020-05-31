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

	var TallerValues []interface{}
	var TallerStrings []string
	TallerValues = utilities.ObjectValues(Taller)
	TallerStrings = utilities.ObjectFields(Taller)

	result, err := utilities.InsertObject("Taller", TallerValues, TallerStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Taller creado")
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

		var TallerValues []interface{}
		var TallerStrings []string
		TallerValues = utilities.ObjectValues(Taller)
		TallerStrings = utilities.ObjectFields(Taller)


		//Limpia de los atributos del objeto
		for i := 0; i < 3 ; i++ {
			if TallerValues[i] == 0 {
				TallerValues[i] = nil
			}
		}
		

		for i := 3; i < len(TallerStrings); i++ {
			if TallerValues[i] == "" {
				TallerValues[i] = nil
			}
		}

		TallerRows, err := utilities.GetObject([]string{"Taller"}, nil, TallerStrings, TallerValues)
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

//QueryToTaller : Metodo que transforma la consulta a objetos Taller
func QueryToTaller(result *sql.Rows) ([]models.Taller, error) {
	var TallerAux models.Taller
	var recipents []models.Taller
	for result.Next() {
		err := result.Scan(&TallerAux.ID, &TallerAux.Nombre, &TallerAux.Enfoque, &TallerAux.IDTipo, &TallerAux.Duracion )
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