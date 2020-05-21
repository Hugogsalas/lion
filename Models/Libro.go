package models

//Libro : Modelo de Libro
type Libro struct {
	ID          int     `json:"id"`
	IDAutor     int     `json:"idAutor"`
	Precio      float64 `json:"precio"`
	Titulo      string  `json:"titulo"`
}
