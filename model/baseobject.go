package model

import (
	"fmt"
)

// ICarbonEmitter interface (Interface Segregation and Dependency Inversion)
type ICarbonEmitter interface {
	CalculateCarbon() float64
}

// Describable interface for the Describe method
type Describable interface {
	Describe() string
}

// BaseObject struct (Single Responsibility)
type BaseObject struct {
	Name               string
	EmbodiedProfile    *CarbonProfile
	OperationalProfile *CarbonProfile
}

// CalculateCarbon method for BaseObject (Liskov Substitution and Open/Closed)
func (bo BaseObject) CalculateCarbon() float64 {
	totalCarbon := 0.0
	if bo.EmbodiedProfile != nil {
		totalCarbon += bo.EmbodiedProfile.EmissionRate
	}
	if bo.OperationalProfile != nil {
		totalCarbon += bo.OperationalProfile.EmissionRate
	}
	return totalCarbon
}

// Describe method for BaseObject (Single Responsibility)
func (bo BaseObject) Describe() string {
	description := fmt.Sprintf("BaseObject: %s\n", bo.Name)
	if bo.EmbodiedProfile != nil {
		description += fmt.Sprintf("  Embodied Profile: Emission Rate = %.2f\n", bo.EmbodiedProfile.EmissionRate)
	}
	if bo.OperationalProfile != nil {
		description += fmt.Sprintf("  Operational Profile: Emission Rate = %.2f\n", bo.OperationalProfile.EmissionRate)
	}
	return description
}
