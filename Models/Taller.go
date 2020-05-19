package models

//Taller : Modelo de Taller
type Taller struct {
	ID      int `json:"id"`
	Nombre  int `json:"nombre"`
	Enfoque int `json:"Enfoque"`
	IDTipo  int `json:"idTipo"`
}
