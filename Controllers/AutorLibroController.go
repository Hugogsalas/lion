package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/mitchellh/mapstructure"

	models "../Models"
	utilities "../Utilities"
)

//CreateAutorLibro : Metodo de insercion de un nuevo AutorLibro
func CreateAutorLibro(writter http.ResponseWriter, request *http.Request) {
	var AutorLibro models.AutorLibro

	err := json.NewDecoder(request.Body).Decode(&AutorLibro)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	AutorLibroStrings, AutorLibroValues := utilities.ObjectFields(AutorLibro)

	result, err := utilities.InsertObject("AutorLibro", AutorLibroValues, AutorLibroStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "AutorLibro creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetAutorLibro : metodo para conseguir la relacion Αutor-libro
func GetAutorLibro(writter http.ResponseWriter, request *http.Request) {
	var AutorLibro models.AutorLibro
	err := json.NewDecoder(request.Body).Decode(&AutorLibro)
	jsonResponse := simplejson.New()

	if err == nil {

		AutorLibroRows, err := utilities.CallStorageProcedure("PAAutorLibro", []interface{}{AutorLibro.IDAutor, AutorLibro.IDLibro})
		if err == nil {
			var AutorLibroResultado []map[string]interface{}

			if AutorLibro.IDAutor == 0 {
				AutorLibroResultado, err = LibrosWithAutores(AutorLibroRows)
			} else {
				AutorLibroResultado, err = AutoresWithLibros(AutorLibroRows)
			}

			if err == nil {
				if len(AutorLibroResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "AutorLibro encontrado")
					jsonResponse.Set("AutorLibro", AutorLibroResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron AutorLibros")
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

//UpdateAutorLibro : Metodo que actualiza AutorLibro segun parametros
func UpdateAutorLibro(writter http.ResponseWriter, request *http.Request) {
	var lastAutorLibro models.AutorLibro
	var newAutorLibro models.AutorLibro
	var recipient map[string]interface{}
	err := json.NewDecoder(request.Body).Decode(&recipient)
	jsonResponse := simplejson.New()
	if err == nil {

		mapstructure.Decode(recipient["filter"], &lastAutorLibro)
		mapstructure.Decode(recipient["update"], &newAutorLibro)

		AutorLibroFilters, AutorLibroFiltersValues := utilities.ObjectFields(lastAutorLibro)
		AutorLibroStrings,AutorLibroValues := utilities.ObjectFields(newAutorLibro)


		AutorLibroRows, err := utilities.UpdateObject("AutorLibro", AutorLibroFilters, AutorLibroFiltersValues, AutorLibroStrings, AutorLibroValues)
		if err == nil {

			if AutorLibroRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "AutorLibro actualizado")

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

//DeleteAutorLibro : Metodo que elimina AutorLibro segun parametros
func DeleteAutorLibro(writter http.ResponseWriter, request *http.Request) {
	var AutorLibro models.AutorLibro
	err := json.NewDecoder(request.Body).Decode(&AutorLibro)
	jsonResponse := simplejson.New()
	if err == nil {

		AutorLibroStrings, AutorLibroValues := utilities.ObjectFields(AutorLibro)

		AutorLibroDel, err := utilities.DeleteObject("AutorLibro", AutorLibroStrings, AutorLibroValues)
		if err == nil {

			if AutorLibroDel {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "AutorLibro eliminado")

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


//LibrosWithAutores : metodo que combierte una consulta a una relacion Libro con autores descritos
func LibrosWithAutores(result *sql.Rows) ([]map[string]interface{}, error) {
	var LibroAux models.Libro
	var AutorAux models.Autor
	var Libros []models.Libro
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&AutorAux.ID,
			&AutorAux.Nombre,
			&AutorAux.ApellidoPaterno,
			&AutorAux.ApellidoMaterno,
			&LibroAux.ID,
			&LibroAux.Titulo,
			&LibroAux.Precio)

		if err != nil {
			return nil, err
		}

		index := utilities.Ιndexof(LibrosToInterfaces(Libros), LibroAux)
		if index == -1 {
			Libros = append(Libros, LibroAux)
			newLibroInfo := map[string]interface{}{
				"id":      LibroAux.ID,
				"titulo":  LibroAux.Titulo,
				"precio":  LibroAux.Precio,
				"Autores": []models.Autor{AutorAux},
			}
			response = append(response, newLibroInfo)
		} else {
			var lastAutors []models.Autor
			lastAutors = response[index]["Autores"].([]models.Autor)
			response[index]["Autores"] = append(lastAutors, AutorAux)
		}
	}
	return response, nil
}

//AutoresWithLibros : metodo que combierte una consulta a una relacion Autor con libros descritos
func AutoresWithLibros(result *sql.Rows) ([]map[string]interface{}, error) {
	var LibroAux models.Libro
	var AutorAux models.Autor
	var Autores []models.Autor
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&AutorAux.ID,
			&AutorAux.Nombre,
			&AutorAux.ApellidoPaterno,
			&AutorAux.ApellidoMaterno,
			&LibroAux.ID,
			&LibroAux.Titulo,
			&LibroAux.Precio)
		if err != nil {
			return nil, err
		}

		index := utilities.Ιndexof(AutoresToInterfaces(Autores), AutorAux)
		if index == -1 {
			Autores = append(Autores, AutorAux)
			newAutorInfo := map[string]interface{}{
				"id":              AutorAux.ID,
				"nombre":          AutorAux.Nombre,
				"apellidoPaterno": AutorAux.ApellidoPaterno,
				"apellidoMaterno": AutorAux.ApellidoMaterno,
				"Libros":          []models.Libro{LibroAux},
			}
			response = append(response, newAutorInfo)
		} else {
			var lastLibros []models.Libro
			lastLibros = response[index]["Libros"].([]models.Libro)
			response[index]["Libros"] = append(lastLibros, LibroAux)
		}
	}
	return response, nil
}
