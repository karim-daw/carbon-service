package model

import "gorm.io/gorm"

type Indicator struct {
	gorm.Model
	MaterialID uint    `gorm:"index;"`
	A1toA5     float64 `gorm:"type:decimal;"`
	B1toB7     float64 `gorm:"type:decimal;"`
	C1toC4     float64 `gorm:"type:decimal;"`
	D          float64 `gorm:"type:decimal;"`
}
