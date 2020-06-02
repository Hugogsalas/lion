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
	var libro models.Libro
	err := json.NewDecoder(request.Body).Decode(&libro)

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var libValues []interface{}
	var libStrings []string
	libValues = utilities.ObjectValues(libro)
	libStrings = utilities.ObjectFields(libro)

	result, err := utilities.InsertObject("libro", libValues, libStrings)
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
	var libro models.Libro
	err := json.NewDecoder(request.Body).Decode(&libro)
	jsonResponse := simplejson.New()
	if err == nil {

		var libValues []interface{}
		var libStrings []string
		libValues = utilities.ObjectValues(libro)
		libStrings = utilities.ObjectFields(libro)

		//Limpia de los atributos del objeto
		if libValues[0] == 0 {
			libValues[0] = nil
		}
		if libValues[1] == 0.0 {
			libValues[1] = nil
		}
		if libValues[2] == "" {
			libValues[2] = nil
		}

		libRows, err := utilities.GetObject([]string{"Libro"}, nil, libStrings, libValues)
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

		var LibroValues []interface{}
		var LibroStrings []string

		LibroValues = utilities.ObjectValues(Libro)
		LibroStrings = utilities.ObjectFields(Libro)

		LibroValues[0] = nil

		if LibroValues[1] == 0.0 {
			LibroValues[1] = nil
		}
		if LibroValues[2] == "" {
			LibroValues[2] = nil
		}
		

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