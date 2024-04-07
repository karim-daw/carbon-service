package model

import "gorm.io/gorm"

type IndicatorConverter interface {
	ConvertValues(isMetric bool, option int) Indicator
}

// carbon footprint of a material with its carbon footprint
// It contains an Indicator which represents the carbon footprint of the material
// for each phase of the LCA
type Indicator struct {
	gorm.Model
	MaterialID uint    `gorm:"index;"`
	IsMetric   bool    `gorm:"type:bool;"`
	A1toA5     float64 `gorm:"type:decimal;"`
	B1toB7     float64 `gorm:"type:decimal;"`
	C1toC4     float64 `gorm:"type:decimal;"`
	A1         float64 `gorm:"type:decimal;"`
	A2         float64 `gorm:"type:decimal;"`
	A3         float64 `gorm:"type:decimal;"`
	A4         float64 `gorm:"type:decimal;"`
	A5         float64 `gorm:"type:decimal;"`
	B1         float64 `gorm:"type:decimal;"`
	B2         float64 `gorm:"type:decimal;"`
	B3         float64 `gorm:"type:decimal;"`
	B4         float64 `gorm:"type:decimal;"`
	B5         float64 `gorm:"type:decimal;"`
	B6         float64 `gorm:"type:decimal;"`
	B7         float64 `gorm:"type:decimal;"`
	C1         float64 `gorm:"type:decimal;"`
	C2         float64 `gorm:"type:decimal;"`
	C3         float64 `gorm:"type:decimal;"`
	C4         float64 `gorm:"type:decimal;"`
	D          float64 `gorm:"type:decimal;"`
}
