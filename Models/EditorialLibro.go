package models

//EditorialLibro : Modelo de EditorialLibro
type EditorialLibro struct {
	ID          int `json:"id"`
	IDEditorial int `json:"idEditorial"`
	IDLibro     int `json:"idlibro"`
}
