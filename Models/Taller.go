package models

//Taller : Modelo de Taller
type Taller struct {
	ID       int    `json:"id"`
	IDTipo   int    `json:"idTipo"`
	Duracion int    `json:"duracion"`
	Nombre   string `json:"nombre"`
	Enfoque  string `json:"Enfoque"`
	
}
