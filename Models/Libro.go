package models

//Libro : Modelo de Libro
type Libro struct {
	ID          int     `json:"id"`
	Titulo      string  `json:"titulo"`
	IDAutor     int     `json:"idAutor"`
	Precio      float64 `json:"precio"`
}
