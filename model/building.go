package model

import (
	"math"

	"gorm.io/gorm"
)

// Ensure Building struct conforms to the CarbonCalculator interface
var _ CarbonCalculator = &Building{}
var _ ByIndicatorCarbonCalculator = &Building{}

type Building struct {
	gorm.Model
	Name                  string      `gorm:"type:string;unique;not null"`
	GFA                   float64     `gorm:"type:float;"`
	FTF                   float64     `gorm:"type:float;not null"`
	GroundFloorArea       float64     `gorm:"type:float;not null"`
	FacadeArea            float64     `gorm:"type:float;"`
	GlazingArea           float64     `gorm:"type:float;"`
	CladdingArea          float64     `gorm:"type:float;"`
	RoofArea              float64     `gorm:"type:float;"`
	WWR                   float64     `gorm:"type:float;not null"`
	AboveGroundFloorCount int         `gorm:"type:int;not null"`
	UnderGroundFloorCount int         `gorm:"type:int;not null"`
	Assemblies            []*Assembly `gorm:"many2many:building_assemblies;"`
}

// calculate gfa of the building
func (b *Building) CalculateGFA() float64 {
	return b.GroundFloorArea * (float64(b.AboveGroundFloorCount) + float64(b.UnderGroundFloorCount))
}

// It calculates the embodied carbon of the building.
func (b *Building) CalculateEmbodiedCarbon() float64 {

	// Calculate the perimeter of the building
	perimeter := calculatePerimeter(2, b.GroundFloorArea)

	// get all the areas
	b.FacadeArea = perimeter * b.FTF * float64(b.AboveGroundFloorCount)
	b.GlazingArea = b.FacadeArea * b.WWR
	b.CladdingArea = (1 - b.WWR) * b.FacadeArea
	b.RoofArea = b.GroundFloorArea

	// Calculate the carbon emissions for each part of the building
	// https://docs.cscale.io/readme/embodied-carbon
	claddingEmission := b.CladdingArea * 8.8 // kgCO2/m2
	glazingEmission := b.FacadeArea * 13.6   // kgCO2/m2
	roofEmission := b.RoofArea * 7.7         // kgCO2/m2

	return claddingEmission + glazingEmission + roofEmission
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

// TODO: #2 needs alot of work
// convert values to metric or imperial and whether its tco2, kgco2, kgco2/m2, kgco2/m2/year
func (b *Building) ConvertValues(isMetric bool, option int) Building {
	for _, assembly := range b.Assemblies {
		assembly.ConvertValues(isMetric, option)
	}
	return *b
}

// calculate perimeter of the building given area and aspect ratio
func calculatePerimeter(aspectRatio, area float64) float64 {

	// error handling
	if aspectRatio <= 0 || area <= 0 {
		panic("aspect ratio and area must be greater than 0")
	}

	// Calculate width and height
	width := math.Sqrt(area * aspectRatio)
	height := area / width

	// Calculate perimeter
	perimeter := 2 * (width + height)
	return perimeter
}
