package model

import (
	"gorm.io/gorm"
)

// Again, ensure Assembly conforms to the interface
var _ CarbonCalculator = &Assembly{}
var _ ByIndicatorCarbonCalculator = &Assembly{}

type Assembly struct {
	gorm.Model
	Name      string      `gorm:"type:string;not null"`
	Buildings []*Building `gorm:"many2many:building_assemblies;"`
	Materials []*Material `gorm:"many2many:assembly_materials;"`
}

func (a Assembly) ComputeWholeLifeCarbon() float64 {
	var totalImpact float64
	for _, material := range a.Materials {
		totalImpact += material.ComputeWholeLifeCarbon()
	}
	return totalImpact
}

func (a Assembly) CalculateCarbonForPhase(phases ...string) float64 {
	var total float64
	for _, material := range a.Materials {
		total += material.CalculateCarbonForPhase(phases...)
	}
	return total
}

func (a *Assembly) ConvertValues(isMetric bool, option int) Assembly {
	for _, material := range a.Materials {
		material.ConvertValues(isMetric, option)
	}
	return *a
}
