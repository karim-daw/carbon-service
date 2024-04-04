package model

import (
	"carbon-service/calculation"

	"gorm.io/gorm"
)

// Ensure Building struct conforms to the CarbonImpactCalculator interface
var _ calculation.CarbonImpactCalculator = &Building{}

type Building struct {
	gorm.Model
	Name       string     `gorm:"type:string;unique;not null"`
	Assemblies []Assembly `gorm:"foreignKey:BuildingID"`
}

// implement the ComputeCarbonImpact() method from the CarbonImpactCalculator interface
func (b Building) ComputeCarbonImpact() float64 {
	var totalImpact float64
	for _, assembly := range b.Assemblies {
		// Ensure Assembly implements ComputeCarbonImpact() and correctly calculates its impact
		totalImpact += assembly.ComputeCarbonImpact()
	}
	return totalImpact
}
