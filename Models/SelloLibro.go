package models

//SelloLibro : Modelo de SelloLibro
type SelloLibro struct {
	ID      int `json:"id"`
	IDLibro int `json:"idLibro"`
	IDSello int `json:"idSello"`
}
