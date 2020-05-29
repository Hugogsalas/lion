package models

//ItinerarioTaller : Modelo de ItinerarioTaller
type ItinerarioTaller struct {
	
	IDItinerario int    `json:"idItinerario"`
	IDTaller     int    `json:"idTaller"`
	Horario      string `json:"horario"`
}
