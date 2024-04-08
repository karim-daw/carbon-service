package model

import (
	"carbon-service/service/calculator"

	"gorm.io/gorm"
)

// Indicator represents the carbon footprint of a material
var _ calculator.CarbonCalculator = &Material{}

// Assuming Indicator is defined somewhere in your model package

// Material represents a building material with its carbon footprint
// It contains an Indicator which represents the carbon footprint of the material
// It also has a many-to-many relationship with Assembly
// The Material struct implements the CarbonImpactCalculator interface
type Material struct {
	gorm.Model
	Name       string
	Indicator  Indicator   `gorm:"foreignKey:MaterialID;constraint:OnDelete:CASCADE;"`
	Assemblies []*Assembly `gorm:"many2many:assembly_materials;"`
}

//TODO: implement material having several indicators

// ComputeCarbonImpact calculates the carbon impact of the material
func (m Material) ComputeWholeLifeCarbon() float64 {
	return m.Indicator.A1toA5 + m.Indicator.B1toB7 + m.Indicator.C1toC4 + m.Indicator.D
}

// CalculateCarbonForPhase calculates the carbon impact of the material for specified phases
// examples:
// material.CalculateCarbonForPhase("A1toA5") -> returns the carbon impact of the material for phases A1 to A5
// material.CalculateCarbonForPhase("A1toA5", "B1toB7") -> returns the carbon impact of the material for phases A1 to A5 and B1 to B7
func (m Material) CalculateCarbonForPhase(phases ...string) float64 {
	var total float64
	for _, phase := range phases {
		switch phase {
		case "A1toA5":
			total += m.Indicator.A1toA5
		case "B1toB7":
			total += m.Indicator.B1toB7
		case "C1toC4":
			total += m.Indicator.C1toC4
		case "D":
			total += m.Indicator.D
		}
	}
	return total
}

// ConvertValues converts the carbon values of the material to metric or imperial
// and whether its tco2, kgco2, kgco2/m2, kgco2/m2/year
func (m *Material) ConvertValues(isMetric bool, option int) Material {
	if isMetric {
		m.Indicator.IsMetric = true
	} else {
		m.Indicator.IsMetric = false
	}

	switch option {
	case 1:
		m.Indicator.A1toA5 = m.Indicator.A1toA5 / 1000
		m.Indicator.B1toB7 = m.Indicator.B1toB7 / 1000
		m.Indicator.C1toC4 = m.Indicator.C1toC4 / 1000
		m.Indicator.D = m.Indicator.D / 1000
	case 2:
		m.Indicator.A1toA5 = m.Indicator.A1toA5 * 1000
		m.Indicator.B1toB7 = m.Indicator.B1toB7 * 1000
		m.Indicator.C1toC4 = m.Indicator.C1toC4 * 1000
		m.Indicator.D = m.Indicator.D * 1000
	case 3:
		m.Indicator.A1toA5 = m.Indicator.A1toA5 / 1000
		m.Indicator.B1toB7 = m.Indicator.B1toB7 / 1000
		m.Indicator.C1toC4 = m.Indicator.C1toC4 / 1000
		m.Indicator.D = m.Indicator.D / 1000
		m.Indicator.A1toA5 = m.Indicator.A1toA5 / 1000
		m.Indicator.B1toB7 = m.Indicator.B1toB7 / 1000
		m.Indicator.C1toC4 = m.Indicator.C1toC4 / 1000
		m.Indicator.D = m.Indicator.D / 1000
	}

	return *m
}
