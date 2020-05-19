package models

//ItinerarioTaller : Modelo de ItinerarioTaller
type ItinerarioTaller struct {
	ID           int    `json:"id"`
	IDItenerario int    `json:"idItinerario"`
	IDTaller     int    `json:"idTaller"`
	Horario      string `json:"horario"`
}
