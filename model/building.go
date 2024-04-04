package model

import (
	"carbon-service/impact"

	"gorm.io/gorm"
)

// Ensure Building struct conforms to the CarbonCalculator interface
var _ impact.CarbonCalculator = &Building{}

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
