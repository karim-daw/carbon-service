package model

import (
	"carbon-service/service/calculator"

	"gorm.io/gorm"
)

// Ensure Building struct conforms to the CarbonCalculator interface
var _ calculator.CarbonCalculator = &Building{}

type Building struct {
	gorm.Model
	Name       string      `gorm:"type:string;unique;not null"`
	Assemblies []*Assembly `gorm:"many2many:building_assemblies;"`
}

// ComputeWholeLifeCarbon calculates the total carbon impact of the building.
func (b *Building) ComputeWholeLifeCarbon() float64 {
	var totalImpact float64
	for _, assembly := range b.Assemblies {
		totalImpact += assembly.ComputeWholeLifeCarbon()
	}
	return totalImpact
}

// CalculateCarbonForPhase calculates the building's carbon impact for specified phases.
func (b *Building) CalculateCarbonForPhase(phases ...string) float64 {
	var total float64
	for _, assembly := range b.Assemblies {
		total += assembly.CalculateCarbonForPhase(phases...)
	}
	return total
}

// convert values to metric or imperial and whether its tco2, kgco2, kgco2/m2, kgco2/m2/year
func (b *Building) ConvertValues(isMetric bool, option int) Building {
	for _, assembly := range b.Assemblies {
		assembly.ConvertValues(isMetric, option)
	}
	return *b
}
