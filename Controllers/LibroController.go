package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateLibro : Metodo de insercion de un nuevo Libro
func CreateLibro(writter http.ResponseWriter, request *http.Request) {
	var Libro models.Libro
	err := json.NewDecoder(request.Body).Decode(&Libro)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	LibroStrings,LibroValues := utilities.ObjectFields(Libro)

	result, err := utilities.InsertObject("Libro", LibroValues, LibroStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Libro añadido")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//GetLibro : Metodo que regresa libros segun parametros
func GetLibro(writter http.ResponseWriter, request *http.Request) {
	var Libro models.Libro
	err := json.NewDecoder(request.Body).Decode(&Libro)
	jsonResponse := simplejson.New()
	if err == nil {

		LibroStrings,LibroValues := utilities.ObjectFields(Libro)

		var LibroQuery models.GetQuery
		
		LibroQuery.Tables=[]string{"Libro"}
		LibroQuery.Selects=nil
		LibroQuery.Params=[][]string{LibroStrings}
		LibroQuery.Values=[][]interface{}{LibroValues}
		LibroQuery.Conditions=nil

		libRows, err := utilities.GetObject(LibroQuery)
		if err == nil {
			librosResultado, err := QueryToLibro(libRows)
			
			if err == nil {
				if len(librosResultado) > 0 {
					jsonResponse.Set("Exito", true)
					jsonResponse.Set("Message", "libros encontrados")
					jsonResponse.Set("Libros", librosResultado)
				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", "No se encontraron libros")
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

//UpdateLibro : Metodo que actualiza Libros segun parametros
func UpdateLibro(writter http.ResponseWriter, request *http.Request) {
	var Libro models.Libro
	err := json.NewDecoder(request.Body).Decode(&Libro)
	jsonResponse := simplejson.New()
	if err == nil {

		var LibroFilters []string
		var LibroFiltersValues []interface{}

		LibroFilters = append(LibroFilters, "ID")
		LibroFiltersValues = append(LibroFiltersValues, Libro.ID)

		LibroStrings,LibroValues := utilities.ObjectFields(Libro)

		LibroRows, err := utilities.UpdateObject("Libro", LibroFilters, LibroFiltersValues, LibroStrings, LibroValues)
		if err == nil {

			if LibroRows {

				jsonResponse.Set("Exito", true)
				jsonResponse.Set("Message", "Libro actualizado")

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

//QueryToLibro : Metodo que transforma la consulta a objetos Libro
func QueryToLibro(result *sql.Rows) ([]models.Libro, error) {
	var libroAux models.Libro
	var recipents []models.Libro
	for result.Next() {
		err := result.Scan(&libroAux.ID, &libroAux.Precio, &libroAux.Titulo)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, libroAux)
	}
	return recipents, nil
}

//LibrosToInterfaces : metodo que transforma un arreglo de libros en interfaces
func LibrosToInterfaces(Libros []models.Libro) []interface{} {
	var arrayInterface []interface{}
	for i:=0;i<len(Libros);i++{
		var libroInterface interface{}
		libroInterface=Libros[i]
		arrayInterface=append(arrayInterface,libroInterface)
	}
	return arrayInterface
}