package model

import (
	"carbon-service/calculation"

	"gorm.io/gorm"
)

// Indicator represents the carbon footprint of a material
var _ calculation.CarbonImpactCalculator = &Material{}

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

func (m Material) ComputeCarbonImpact() float64 {
	return m.Indicator.A1toA5 + m.Indicator.B1toB7 + m.Indicator.C1toC4 + m.Indicator.D
}
