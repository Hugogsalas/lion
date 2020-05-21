package models

//Autor : Modelo de Autor
type Autor struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	ApellidoPaterno string `json:"apellidoPaterno"`
	ApellidoMaterno string `json:"apellidoMaterno"`
}
