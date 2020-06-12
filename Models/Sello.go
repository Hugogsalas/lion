package models

//Sello : Modelo de Sello
type Sello struct {
	ID          int    `json:"id"`
	IDEditorial int    `json:"idEditorial"`
	Descripcion string `json:"descripcion"`
}
