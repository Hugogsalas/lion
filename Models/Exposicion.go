package models

//Exposicion : Modelo de Exposicion
type Exposicion struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo"`
	Presentador string `json:"presentador"`
	Duracion    int    `json:"duracion"`
	IDTipo      int    `json:"idTipo"`
}
