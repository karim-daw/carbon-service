package calculator

// CarbonCalculator defines the interface for calculating carbon impact
type CarbonCalculator interface {

	// ComputeWholeLifeCarbon calculates the total carbon impact of a material, assembly, or building
	ComputeWholeLifeCarbon() float64

	// CalculateCarbonForPhase calculates the carbon impact of a material, assembly, or building for a specific phase
	CalculateCarbonForPhase(phase ...string) float64
}

type PhaseBasedCarbonCalculator interface {
	CalculateCarbonForPhase(phase ...string) float64
}
