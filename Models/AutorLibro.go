package models

//AutorLibro : Modelo de AutorLibro
type AutorLibro struct {
	ID          int     `json:"id"`
	IDAutor     int     `json:"idAutor"`
	IDLibro 	int     `json:"idlibro"`
}
