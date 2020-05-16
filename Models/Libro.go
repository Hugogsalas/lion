package models

//Libro : Modelo de Libro
type Libro struct {
	ID        int       `json:"id"`
	Titulo    string    `json:"titulo"`
	Autor     Autor     `json:"autor"`
	Editorial Editorial `json:"editorial"`
	Precio    float64   `json:"precio"`
}
