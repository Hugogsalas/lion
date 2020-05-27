package models

//ItinerarioExposicion : Modelo de ItinerarioExposicion
type ItinerarioExposicion struct {
	
	IDItinerario int    `json:"idItinerario"`
	IDExposicion int    `json:"idExposicion"`
	Horario      string `json:"horario"`
}
