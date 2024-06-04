package model

import "fmt"

type Scenario struct {
	Name      string
	Buildings []*Building
}

// CalculateCarbon method for Scenario (Liskov Substitution and Open/Closed)
func (s Scenario) CalculateCarbon() float64 {
	totalCarbon := 0.0
	for _, building := range s.Buildings {
		totalCarbon += building.CalculateCarbon()
	}
	return totalCarbon
}

// Describe method for Scenario (Single Responsibility)
func (s Scenario) Describe() string {
	description := fmt.Sprintf("Scenario: %s\n", s.Name)
	for _, building := range s.Buildings {
		description += building.Describe()
	}
	return description
}
