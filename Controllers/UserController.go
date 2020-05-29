package controllers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bitly/go-simplejson"

	models "../Models"
	utilities "../Utilities"
)

//CreateUser : Metodo de insercion de un nuevo usuario
func CreateUser(writter http.ResponseWriter, request *http.Request) {
	var usuario models.Usuario
	hash := sha256.New()
	err := json.NewDecoder(request.Body).Decode(&usuario)

	hash.Write([]byte(usuario.Clave))
	usuario.Clave = fmt.Sprintf("%x", hash.Sum(nil))

	json := simplejson.New()
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	var userValues []interface{}
	var useStrings []string
	userValues = utilities.ObjectValues(usuario)
	useStrings = utilities.ObjectFields(usuario)

	result, err := utilities.InsertObject("Usuarios", userValues, useStrings)
	if err != nil {
		json.Set("Exito", false)
		json.Set("Message", err.Error())
	}

	if result {
		json.Set("Exito", true)
		json.Set("Message", "Usuario Creado")
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//LoginUser : Metodo que retorna un usaurio por su correo y clave
func LoginUser(writter http.ResponseWriter, request *http.Request) {
	var usuario models.Usuario
	hash := sha256.New()
	err := json.NewDecoder(request.Body).Decode(&usuario)
	jsonResponse := simplejson.New()

	if err == nil {
		if usuario.Clave != "" {
			if usuario.Correo != "" {
				hash.Write([]byte(usuario.Clave))
				usuario.Clave = fmt.Sprintf("%x", hash.Sum(nil))

				var userValues []interface{}
				var userStrings []string

				userStrings = append(userStrings, "Correo")
				userStrings = append(userStrings, "Clave")

				userValues = append(userValues, usuario.Correo)
				userValues = append(userValues, usuario.Clave)

				UserRow, err := utilities.GetObject([]string{"Usuarios"}, nil, userStrings, userValues)

				if err == nil {
					UserList, err := QueryToUser(UserRow)

					if len(UserList)>0{
						if err == nil {
							jsonResponse.Set("Exito", true)
							jsonResponse.Set("Message", "Usuario obtenido")
							jsonResponse.Set("Usuario", UserList[0])
	
						} else {
							jsonResponse.Set("Exito", false)
							jsonResponse.Set("Message", err.Error())
	
						}
						
					}else{
						jsonResponse.Set("Exito", false)
						jsonResponse.Set("Message", "Correo o clave erroneo")
					}

				} else {
					jsonResponse.Set("Exito", false)
					jsonResponse.Set("Message", err.Error())
				}

			} else {
				jsonResponse.Set("Exito", false)
				jsonResponse.Set("Message", "Falta el campo correo")
			}

		} else {
			jsonResponse.Set("Exito", false)
			jsonResponse.Set("Message", "Falta el campo clave")
		}
	} else {
		jsonResponse.Set("Exito", false)
		jsonResponse.Set("Message", err.Error())
	}

	payload, err := jsonResponse.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}
	writter.Header().Set("Content-Type", "application/json")
	writter.Write(payload)
	return
}

//QueryToUser : Metodo que transforma la consulta a objetos Usuario
func QueryToUser(result *sql.Rows) ([]models.Usuario, error) {
	var userAux models.Usuario
	var recipents []models.Usuario
	for result.Next() {
		err := result.Scan(&userAux.ID, &userAux.Correo, &userAux.Nombre, &userAux.ApellidoPaterno, &userAux.ApellidoMaterno, &userAux.Sexo, &userAux.Clave)
		if err != nil {
			return nil, err
		}
		recipents = append(recipents, userAux)
	}
	return recipents, nil
}
