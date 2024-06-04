package model

import "fmt"

// SystemType enum for system type safety
type SystemType string

const (
	StructureAboveGround       SystemType = "StructureAboveGround"
	StructureBelowGround       SystemType = "StructureBelowGround"
	FacadeGlazed               SystemType = "FacadeGlazed"
	FacadeUnglazed             SystemType = "FacadeUnglazed"
	Roof                       SystemType = "Roof"
	Furniture                  SystemType = "Furniture"
	InteriorPartitionsGlazed   SystemType = "InteriorPartitionsGlazed"
	InteriorPartitionsUnglazed SystemType = "InteriorPartitionsUnglazed"
)

// ISystem interface (Interface Segregation and Dependency Inversion)
type ISystem interface {
	CalculateCarbon(units float64) float64
	Describe() string
}

// System struct
type System struct {
	Type         SystemType
	EmissionRate float64 // kg/co2e per unit
	UnitType     string  // "building" or "assembly"
}

// CalculateCarbon method for System
func (s System) CalculateCarbon(units float64) float64 {
	return s.EmissionRate * units
}

// Describe method for System
func (s System) Describe() string {
	return fmt.Sprintf("System: %s, Emission Rate: %.2f kg/co2e per %s", s.Type, s.EmissionRate, s.UnitType)
}
