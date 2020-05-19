package models

//Autor : Modelo de Autor
type Autor struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	ApellidoPaterno string `json:"apellidoPaterno"`
	ApellinaMaterno string `json:"apellinaMaterno"`
}
