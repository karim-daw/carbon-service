package model

import "gorm.io/gorm"

type Indicator interface {
	A1toA5() float64
	B1toB7() float64
	C1toC4() float64
	// array for all phases from A1 to D
	GetIndicators() []float64
}

type IndicatorConverter interface {
	ConvertValues(isMetric bool, option int)
}

// carbon footprint of a material with its carbon footprint
// It contains an Indicator which represents the carbon footprint of the material
// for each phase of the LCA

var _ Indicator = &Gwp{}
var _ IndicatorConverter = &Gwp{}

type Gwp struct {
	gorm.Model
	MaterialID uint    `gorm:"index;"`
	IsMetric   bool    `gorm:"type:bool;"`
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

// Returns the sum of all phases from A1 to A5
func (g Gwp) A1toA5() float64 {
	return g.A1 + g.A2 + g.A3 + g.A4 + g.A5
}

// Returns the sum of all phases from B1 to B7
func (g Gwp) B1toB7() float64 {
	return g.B1 + g.B2 + g.B3 + g.B4 + g.B5 + g.B6 + g.B7
}

// Returns the sum of all phases from C1 to C4
func (g Gwp) C1toC4() float64 {
	return g.C1 + g.C2 + g.C3 + g.C4
}

// Returns an array of all phases from A1 to D
func (g Gwp) GetIndicators() []float64 {
	return []float64{g.A1, g.A2, g.A3, g.A4, g.A5, g.B1, g.B2, g.B3, g.B4, g.B5, g.B6, g.B7, g.C1, g.C2, g.C3, g.C4, g.D}
}

// ConvertValues converts the carbon values of the material to metric or imperial
// and whether its tco2, kgco2, kgco2/m2, kgco2/m2/year
func (g *Gwp) ConvertValues(isMetric bool, option int) {
	if isMetric {
		g.IsMetric = true
	} else {
		g.IsMetric = false
	}

	switch option {
	case 1:
		g.A1 = g.A1 / 1000
		g.A2 = g.A2 / 1000
		g.A3 = g.A3 / 1000
		g.A4 = g.A4 / 1000
		g.A5 = g.A5 / 1000
		g.B1 = g.B1 / 1000
		g.B2 = g.B2 / 1000
		g.B3 = g.B3 / 1000
		g.B4 = g.B4 / 1000
		g.B5 = g.B5 / 1000
		g.B6 = g.B6 / 1000
		g.B7 = g.B7 / 1000
		g.C1 = g.C1 / 1000
		g.C2 = g.C2 / 1000
		g.C3 = g.C3 / 1000
		g.C4 = g.C4 / 1000
		g.D = g.D / 1000
	case 2:
		g.A1 = g.A1 * 1000
		g.A2 = g.A2 * 1000
		g.A3 = g.A3 * 1000
		g.A4 = g.A4 * 1000
		g.A5 = g.A5 * 1000
		g.B1 = g.B1 * 1000
		g.B2 = g.B2 * 1000
		g.B3 = g.B3 * 1000
		g.B4 = g.B4 * 1000
		g.B5 = g.B5 * 1000
		g.B6 = g.B6 * 1000
		g.B7 = g.B7 * 1000
		g.C1 = g.C1 * 1000
		g.C2 = g.C2 * 1000
		g.C3 = g.C3 * 1000
		g.C4 = g.C4 * 1000
		g.D = g.D * 1000
	case 3:
		g.A1 = g.A1 / 1000
		g.A2 = g.A2 / 1000
		g.A3 = g.A3 / 1000
		g.A4 = g.A4 / 1000
		g.A5 = g.A5 / 1000
		g.B1 = g.B1 / 1000
		g.B2 = g.B2 / 1000
		g.B3 = g.B3 / 1000
		g.B4 = g.B4 / 1000
		g.B5 = g.B5 / 1000
		g.B6 = g.B6 / 1000
		g.B7 = g.B7 / 1000
		g.C1 = g.C1 / 1000
		g.C2 = g.C2 / 1000
		g.C3 = g.C3 / 1000
		g.C4 = g.C4 / 1000
		g.D = g.D / 1000
	}
}
