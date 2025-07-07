package models

import "time"

// structures for DB
type Car struct {
	ID               int       `json:"id"`
	ManufacturerID   int       `json:"manufacturerId"`
	Brand            string    `json:"brand"`
	Model            string    `json:"model"`
	CreatedYear      int       `json:"createdYear"`
	EngineCapacity   float64   `json:"engineCapacity"`
	Power            int       `json:"power"`
	DriveType        string    `json:"driveType"`
	TransmissionType string    `json:"transmissionType"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
