package controllers

import (
	"encoding/json"
	"net/http"
	"database/sql"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateSelloLibro : Metodo de insercion de un nuevo SelloLibro
func CreateSelloLibro(writter http.ResponseWriter, request *http.Request) {
	var SelloLibro models.SelloLibro
	
	err := json.NewDecoder(request.Body).Decode(&SelloLibro)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var SelloLibroValues []interface{}
	var SelloLibroStrings []string
	SelloLibroValues = utilities.ObjectValues(SelloLibro)
	SelloLibroStrings = utilities.ObjectFields(SelloLibro)

	result, err := utilities.InsertObject("SelloLibro", SelloLibroValues, SelloLibroStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "SelloLibro creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetSelloLibro : metodo que retorna  una relacion Sello-Libro
func GetSelloLibro(writter http.ResponseWriter, request *http.Request) {
	var SelloLibro models.SelloLibro
	err := json.NewDecoder(request.Body).Decode(&SelloLibro)
	jsonResponse := simplejson.New()

	if err == nil {

		SelloLibroRows, err := utilities.CallStorageProcedure("PASelloLibro", []interface{}{SelloLibro.IDSello, SelloLibro.IDLibro})
		if err == nil {
			var SelloLibroResultado []map[string]interface{}

			if SelloLibro.IDLibro == 0 {
				SelloLibroResultado, err = SelloWithLibros(SelloLibroRows)
			} else {
				SelloLibroResultado, err = LibroWithSellos(SelloLibroRows)
			}

			if err == nil {
				if len(SelloLibroResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "AutorLibro encontrado")
					jsonResponse.Set("SelloLibro", SelloLibroResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron SelloLibros")
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

//SelloWithLibros : metodo que combierte una consulta a una relacion Sello con libros descritos
func SelloWithLibros(result *sql.Rows) ([]map[string]interface{}, error) {
	var LibroAux models.Libro
	var SelloAux models.Sello
	var Sellos []models.Sello
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&SelloAux.ID,
			&SelloAux.Descripcion,
			&LibroAux.ID,
			&LibroAux.Titulo,
			&LibroAux.Precio)
		if err != nil {
			return nil, err
		}
			
		index:=utilities.Ιndexof(SellosToInterfaces(Sellos),SelloAux)
		if index==-1{
			Sellos=append(Sellos,SelloAux)
			newAutorInfo:=map[string]interface{}{
				"id":SelloAux.ID,
				"descripcion":SelloAux.Descripcion,
				"Libros":[]models.Libro{LibroAux},
			}
			response=append(response,newAutorInfo)
		}else{
			var Libros []models.Libro
			Libros=response[index]["Libros"].([]models.Libro)
			response[index]["Libros"]=append(Libros,LibroAux)
		}
	}
	return response, nil
}

//LibroWithSellos : metodo que combierte una consulta a una relacion libros con Sellos descritos
func LibroWithSellos(result *sql.Rows) ([]map[string]interface{}, error) {
	var LibroAux models.Libro
	var SelloAux models.Sello
	var Libros []models.Libro
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&SelloAux.ID,
			&SelloAux.Descripcion,
			&LibroAux.ID,
			&LibroAux.Titulo,
			&LibroAux.Precio)
		if err != nil {
			return nil, err
		}
			
		index:=utilities.Ιndexof(LibrosToInterfaces(Libros),LibroAux)
		if index==-1{
			Libros=append(Libros,LibroAux)
			newAutorInfo:=map[string]interface{}{
				"id":LibroAux.ID,
				"titulo":LibroAux.Titulo,
				"precio":LibroAux.Precio,
				"Sellos":[]models.Sello{SelloAux},
			}
			response=append(response,newAutorInfo)
		}else{
			var sellos []models.Sello
			sellos=response[index]["Sellos"].([]models.Sello)
			response[index]["Sellos"]=append(sellos,SelloAux)
		}
	}
	return response, nil
}