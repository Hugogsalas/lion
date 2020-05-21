package models

//Usuario : Modelo de Usuario
type Usuario struct {
	ID				int    `json:"id"`
	Correo          string `json:"correo"`
	Nombre          string `json:"nombre"`
	ApellidoPaterno string `json:"apellidoPaterno"`
	ApellidoMaterno string `json:"apellidoMaterno"`
	Sexo          	string `json:"sexo"`
	Clave			string `json:"clave"`
}
