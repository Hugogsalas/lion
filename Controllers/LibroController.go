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


//GetLib : Metodo que regresa autores segun parametros
func GetLib(writter http.ResponseWriter, request *http.Request) {
	var libro models.Libro
	err := json.NewDecoder(request.Body).Decode(&libro)
	jsonResponse := simplejson.New()
	if err == nil {

		var libValues []interface{}
		var libStrings []string
		libValues = utilities.ObjectValues(libro)
		libStrings = utilities.ObjectFields(libro)

		fmt.Println(libStrings)
		fmt.Println(libValues)

		//Limpia de los atributos del objeto
		for i := 0; i < 3; i++ {
			if libValues[i] == 0{
				libValues[i] = nil
			}
		}
		if libValues[3] == "" {
			libValues[3] = nil
		}


		libRows, err := utilities.GetObject("libro", nil, libStrings, libValues)
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

//QueryToLibro : Metodo que transforma la consulta a objetos Libro
func QueryToLibro(result *sql.Rows) ([]models.Libro, error) {
	var libroAux models.Libro
	var recipents []models.Libro
	for result.Next() {
		err := result.Scan(&libroAux.ID, &libroAux.IDAutor, &libroAux.Precio, &libroAux.Titulo)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, libroAux)
	}
	return recipents, nil
}