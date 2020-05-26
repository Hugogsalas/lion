package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bitly/go-simplejson"

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

	var AutorLibroValues []interface{}
	var AutorLibroStrings []string
	AutorLibroValues = utilities.ObjectValues(AutorLibro)
	AutorLibroStrings = utilities.ObjectFields(AutorLibro)

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
		
		AutorLibroRows, err := utilities.CallStorageProcedure("PAAutorLibro",[]interface{}{AutorLibro.IDAutor,AutorLibro.IDLibro})
		if err == nil {
			AutorLibroResultado, err := QueryToAutorLibro(AutorLibroRows)
			
			/*if AutorLibroValues[0]==nil{
				AutorLibroResponse,err := GetRelations(AutorLibroResultado)
			}else{
				
			}*/

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

//QueryToAutorLibro : metodo que combierte una consulta a una relacion autor-libro
func QueryToAutorLibro(result *sql.Rows) ([]models.AutorLibro, error) {
	var ΑutorLibroAux models.AutorLibro
	var recipents []models.AutorLibro
	for result.Next() {
		err := result.Scan(&ΑutorLibroAux.IDAutor, &ΑutorLibroAux.IDLibro)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, ΑutorLibroAux)
	}
	return recipents, nil
}

//GetRelations : metodo que trae los autores y libros relacionados
/*func GetRelations(relations []models.AutorLibro) ([]map[string]interface{}, error) {
	var relationsParse []map[string]interface{}
	var newRelation map[string]interface{}
	var idsAutorEncontrados []interface{}
	var idsLibrosEncontrados []interface{}
	var Autores []models.Autor
	var Libros []models.Libro
	for i := 0; i < len(relations); i++ {
		rel := relations[i]
		indexAutor := utilities.Ιndexof(idsAutorEncontrados, rel.IDAutor)
		if indexAutor == -1 {
			searchAutor, err := utilities.GetObject("Autor", nil, []string{"ID"}, []interface{}{rel.IDAutor})
			if err == nil {
				newΑutor, err := QueryToAutor(searchAutor)
				if err==nil{
					Autores=append(Autores,newΑutor[0])
					newRelation["Autor"]
				}else{
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {

		}
		indexLibro := utilities.Ιndexof(idsLibrosEncontrados, rel.IDLibro)
		if indexLibro == -1 {

		} else {

		}
	}
	return relationsParse, nil
}
*/