package model

import "gorm.io/gorm"

type Building struct {
	gorm.Model         // Embedded, providing ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string  `json:"name"`
	Location   string  `json:"location"`
	EC         float64 `json:"ec"` // Example: Environmental Compliance
	OC         float64 `json:"oc"` // Example: Operational Cost
}
