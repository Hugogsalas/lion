package controllers

import (
	"encoding/json"
	"net/http"
	"database/sql"

	"github.com/bitly/go-simplejson"

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

	var SalaExposicionValues []interface{}
	var SalaExposicionestrings []string
	SalaExposicionValues = utilities.ObjectValues(SalaExposicion)
	SalaExposicionestrings = utilities.ObjectFields(SalaExposicion)

	result, err := utilities.InsertObject("SalaExposicion", SalaExposicionValues, SalaExposicionestrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Sala-Exposicion creado")
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
					jsonResponse.Set("Message", "AutorExposicion encontrado")
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

//SalaWithExposiciones : metodo que combierte una consulta a una relacion Sala con Exposiciones descritos
func SalaWithExposiciones(result *sql.Rows) ([]map[string]interface{}, error) {
	var ExposicionAux models.Exposicion
	var SalaAux models.Sala
	var Salas []models.Sala
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&ExposicionAux.ID,
			&ExposicionAux.Presentador,
			&ExposicionAux.Titulo,
			&ExposicionAux.Duracion,
			&SalaAux.ID,
			&SalaAux.Nombre)

		if err != nil {
			return nil, err
		}
		
			
		index:=utilities.Ιndexof(SalasToInterfaces(Salas),SalaAux)
		if index==-1{
			Salas=append(Salas,SalaAux)
			newExposicionInfo:=map[string]interface{}{
				"id":SalaAux.ID,
				"nombre":SalaAux.Nombre,
				"Exposiciones":[]models.Exposicion{ExposicionAux},
			}
			response=append(response,newExposicionInfo)
		}else{
			var lastExposiciones []models.Exposicion
			lastExposiciones=response[index]["Exposiciones"].([]models.Exposicion)
			response[index]["Exposiciones"]=append(lastExposiciones,ExposicionAux)
		}
	}
	return response, nil
}

//ExposicionesWithSala : metodo que combierte una consulta a una relacion libors con Salas descritos
func ExposicionesWithSala(result *sql.Rows) ([]map[string]interface{}, error) {
	var ExposicionAux models.Exposicion
	var SalaAux models.Sala
	var Exposiciones []models.Exposicion
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&ExposicionAux.ID,
			&ExposicionAux.Presentador,
			&ExposicionAux.Titulo,
			&ExposicionAux.Duracion,
			&SalaAux.ID,
			&SalaAux.Nombre)
		if err != nil {
			return nil, err
		}
			
		index:=utilities.Ιndexof(ExposicionesToInterfaces(Exposiciones),ExposicionAux)
		if index==-1{
			Exposiciones=append(Exposiciones,ExposicionAux)
			newAutorInfo:=map[string]interface{}{
				"id":ExposicionAux.ID,
				"Duracion":ExposicionAux.Duracion,
				"Titulo":ExposicionAux.Titulo,
				"Presentador":ExposicionAux.Presentador,
				"Salas":[]models.Sala{SalaAux},
			}
			response=append(response,newAutorInfo)
		}else{
			var Salas []models.Sala
			Salas=response[index]["Salas"].([]models.Sala)
			response[index]["Salas"]=append(Salas,SalaAux)
		}
	}
	return response, nil
}