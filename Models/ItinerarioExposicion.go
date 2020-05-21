package models

//ItinerarioExposicion : Modelo de ItinerarioExposicion
type ItinerarioExposicion struct {
	
	IDItenerario int    `json:"idItinerario"`
	IDExposicion int    `json:"idExposicion"`
	Horario      string `json:"horario"`
}
