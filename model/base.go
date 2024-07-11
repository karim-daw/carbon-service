package model

// CarbonCalculator defines the interface for calculating carbon impact
type CarbonCalculator interface {
	ComputeWholeLifeCarbon() float64
}

type ByIndicatorCarbonCalculator interface {
	CalculateCarbonForPhase(phase ...string) float64
}

// EmbodiedCarbonCalculator defines the interface for calculating embodied carbon
type EmbodiedCarbonCalculator interface {
	CalculateEmbodiedCarbon() float64
}

// OperationalCarbonCalculator defines the interface for calculating operational carbon
type OperationalCarbonCalculator interface {
	CalculateOperationalCarbon() float64
}
