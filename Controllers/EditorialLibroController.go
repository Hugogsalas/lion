package controllers

import (
	"encoding/json"
	"net/http"
	"database/sql"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateEditorialLibro : Metodo de insercion de un nuevo EditorialLibro
func CreateEditorialLibro(writter http.ResponseWriter, request *http.Request) {
	var EditorialLibro models.EditorialLibro
	
	err := json.NewDecoder(request.Body).Decode(&EditorialLibro)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var EditorialLibroValues []interface{}
	var EditorialLibroStrings []string
	EditorialLibroValues = utilities.ObjectValues(EditorialLibro)
	EditorialLibroStrings = utilities.ObjectFields(EditorialLibro)

	result, err := utilities.InsertObject("EditorialLibro", EditorialLibroValues, EditorialLibroStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "EditorialLibro creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetEditorialLibro : metodo que retorna  una relacion Editorial-Libro
func GetEditorialLibro(writter http.ResponseWriter, request *http.Request) {
	var EditorialLibro models.EditorialLibro
	err := json.NewDecoder(request.Body).Decode(&EditorialLibro)
	jsonResponse := simplejson.New()

	if err == nil {

		EditorialLibroRows, err := utilities.CallStorageProcedure("PAEditorialLibro", []interface{}{EditorialLibro.IDEditorial, EditorialLibro.IDLibro})
		if err == nil {
			var EditorialLibroResultado []map[string]interface{}

			if EditorialLibro.IDLibro == 0 {
				EditorialLibroResultado, err = EditorialWithLibros(EditorialLibroRows)
			} else {
				EditorialLibroResultado, err = LibrosWithEditorial(EditorialLibroRows)
			}

			if err == nil {
				if len(EditorialLibroResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "AutorLibro encontrado")
					jsonResponse.Set("EditorialLibro", EditorialLibroResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron EditorialLibros")
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

//EditorialWithLibros : metodo que combierte una consulta a una relacion Editorial con libros descritos
func EditorialWithLibros(result *sql.Rows) ([]map[string]interface{}, error) {
	var LibroAux models.Libro
	var EditorialAux models.Editorial
	var Editoriales []models.Editorial
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&LibroAux.ID,
			&LibroAux.Titulo,
			&LibroAux.Precio,
			&EditorialAux.ID,
			&EditorialAux.Nombre)

		if err != nil {
			return nil, err
		}
			
		index:=utilities.Ιndexof(EditorialesToInterfaces(Editoriales),EditorialAux)
		if index==-1{
			Editoriales=append(Editoriales,EditorialAux)
			newLibroInfo:=map[string]interface{}{
				"id":EditorialAux.ID,
				"nombre":EditorialAux.Nombre,
				"Libros":[]models.Libro{LibroAux},
			}
			response=append(response,newLibroInfo)
		}else{
			var lastLibros []models.Libro
			lastLibros=response[index]["Libros"].([]models.Libro)
			response[index]["Libros"]=append(lastLibros,LibroAux)
		}
	}
	return response, nil
}

//LibrosWithEditorial : metodo que combierte una consulta a una relacion libors con editoriales descritos
func LibrosWithEditorial(result *sql.Rows) ([]map[string]interface{}, error) {
	var LibroAux models.Libro
	var EditorialAux models.Editorial
	var Libros []models.Libro
	var response []map[string]interface{}
	for result.Next() {
		err := result.Scan(
			&LibroAux.ID,
			&LibroAux.Titulo,
			&LibroAux.Precio,
			&EditorialAux.ID,
			&EditorialAux.Nombre)
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
				"Editoriales":[]models.Editorial{EditorialAux},
			}
			response=append(response,newAutorInfo)
		}else{
			var Editoriales []models.Editorial
			Editoriales=response[index]["Editoriales"].([]models.Editorial)
			response[index]["Editoriales"]=append(Editoriales,EditorialAux)
		}
	}
	return response, nil
}