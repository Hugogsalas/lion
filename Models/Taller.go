package models

//Taller : Modelo de Taller
type Taller struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Enfoque  string `json:"Enfoque"`
	IDTipo   int    `json:"idTipo"`
	Duracion int    `json:"duracion"`
}
