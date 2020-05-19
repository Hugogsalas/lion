package models

//ItinerarioExposicion : Modelo de ItinerarioExposicion
type ItinerarioExposicion struct {
	ID           int    `json:"id"`
	IDItenerario int    `json:"idItinerario"`
	IDExposicion int    `json:"idExposicion"`
	Horario      string `json:"horario"`
}
