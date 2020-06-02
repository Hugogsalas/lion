package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateStan : Metodo de insercion de un nuevo Stan
func CreateStan(writter http.ResponseWriter, request *http.Request) {
	var Stan models.Stan

	err := json.NewDecoder(request.Body).Decode(&Stan)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var StanValues []interface{}
	var StanStrings []string
	StanValues = utilities.ObjectValues(Stan)
	StanStrings = utilities.ObjectFields(Stan)

	result, err := utilities.InsertObject("Stan", StanValues, StanStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Stan creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetStan : Metodo que regresa Stanes segun parametros
func GetStan(writter http.ResponseWriter, request *http.Request) {
	var Stan models.Stan
	err := json.NewDecoder(request.Body).Decode(&Stan)
	jsonResponse := simplejson.New()
	if err == nil {

		var StanValues []interface{}
		var StanStrings []string
		StanValues = utilities.ObjectValues(Stan)
		StanStrings = utilities.ObjectFields(Stan)

		//Limpia de los atributos del objeto

		for i := 0; i < 3; i++ {
			if StanValues[i] == 0 {
				StanValues[i] = nil
			}
		}

		StanRows, err := utilities.GetObject([]string{"Stan"}, nil, StanStrings, StanValues)
		if err == nil {
			StanesResultado, err := QueryToStan(StanRows)
			if err == nil {
				if len(StanesResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "Stans encontrados")
					jsonResponse.Set("Stans", StanesResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron Stans")
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

//GetStanWithEditorial : Metodo que regresa Stanes con el desgloce de la editorial segun parametros
func GetStanWithEditorial(writter http.ResponseWriter, request *http.Request) {
	var Stan models.Stan
	err := json.NewDecoder(request.Body).Decode(&Stan)
	jsonResponse := simplejson.New()
	if err == nil {

		var StanValues []interface{}
		StanValues = utilities.ObjectValues(Stan)

		StanRows, err := utilities.CallStorageProcedure("PAStanEditoriales", StanValues)

		var StanesResultado []map[string]interface{}

		if err == nil {
			if Stan.ID != 0 || Stan.Numero != 0{
				StanesResultado, err = StanWithEditorial(StanRows)
			} else {
				StanesResultado, err = EditorialwithStans(StanRows)
			}
			if err == nil {
				if len(StanesResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "Stans encontrados")
					jsonResponse.Set("Stans", StanesResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron Stans")
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

//UpdateStan : Metodo que actualiza Stans segun parametros
func UpdateStan(writter http.ResponseWriter, request *http.Request) {
	var Stan models.Stan
	err := json.NewDecoder(request.Body).Decode(&Stan)
	jsonResponse := simplejson.New()
	if err == nil {

		var StanFilters []string
		var StanFiltersValues []interface{}

		StanFilters = append(StanFilters, "ID")
		StanFiltersValues = append(StanFiltersValues, Stan.ID)

		var StanValues []interface{}
		var StanStrings []string

		StanValues = utilities.ObjectValues(Stan)
		StanStrings = utilities.ObjectFields(Stan)

		StanValues[0] = nil

		for i := 1; i < 3; i++ {
			if StanValues[i] == 0 {
				StanValues[i] = nil
			}
		}

		StanRows, err := utilities.UpdateObject("Stan", StanFilters, StanFiltersValues, StanStrings, StanValues)
		if err == nil {

			if StanRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "Stan actualizado")

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

//QueryToStan : Metodo que transforma la consulta a objetos Stan
func QueryToStan(result *sql.Rows) ([]models.Stan, error) {
	var StanAux models.Stan
	var recipents []models.Stan
	for result.Next() {
		err := result.Scan(&StanAux.ID, &StanAux.IDEditorial, &StanAux.Numero)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, StanAux)
	}
	return recipents, nil
}

//EditorialwithStans : Metodo que transforma la consulta a objetos Editorial con sus stans descrita
func EditorialwithStans(result *sql.Rows) ([]map[string]interface{}, error) {
	var EditorialAux models.Editorial
	var StanAux models.Stan
	var Editoriales []models.Editorial
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&StanAux.ID,
			&StanAux.Numero,
			&EditorialAux.ID,
			&EditorialAux.Nombre)
		if err != nil {
			return nil, err
		}

		index := utilities.Ιndexof(EditorialesToInterfaces(Editoriales), EditorialAux)
		if index == -1 {
			Editoriales = append(Editoriales, EditorialAux)
			newEditorialInfo := map[string]interface{}{
				"id":     EditorialAux.ID,
				"nombre": EditorialAux.Nombre,
				"Stans": []map[string]interface{}{map[string]interface{}{
					"id":     StanAux.ID,
					"numero": StanAux.Numero,
				}},
			}
			response = append(response, newEditorialInfo)
		} else {
			var Stans []map[string]interface{}
			Stans = response[index]["Stans"].([]map[string]interface{})
			response[index]["Stans"] = append(Stans, map[string]interface{}{
				"id":     StanAux.ID,
				"numero": StanAux.Numero,
			})
		}
	}
	return response, nil
}

//StanWithEditorial : Metodo que transforma la consulta a objetos Stan con editorial descrita
func StanWithEditorial(result *sql.Rows) ([]map[string]interface{}, error) {
	var EditorialAux models.Editorial
	var StanAux models.Stan
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&StanAux.ID,
			&StanAux.Numero,
			&EditorialAux.ID,
			&EditorialAux.Nombre)

		if err != nil {
			return nil, err
		}

		newStanInfo := map[string]interface{}{
			"id":     StanAux.ID,
			"numero": StanAux.Numero,
			"Editorial": map[string]interface{}{
				"id":     EditorialAux.ID,
				"nombre": EditorialAux.Nombre,
			},
		}
		response = append(response, newStanInfo)
	}
	return response, nil
}

//StanesToInterfaces : metodo que transforma un arreglo de Stanes en interfaces
func StanesToInterfaces(Stanes []models.Stan) []interface{} {
	var arrayInterface []interface{}
	for i := 0; i < len(Stanes); i++ {
		var StanInterface interface{}
		StanInterface = Stanes[i]
		arrayInterface = append(arrayInterface, StanInterface)
	}
	return arrayInterface
}
