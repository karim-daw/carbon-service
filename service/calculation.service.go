// service/calculation.go

package service

import (
	"carbon-service/service/calculator"
	"sync"
)

type CalculationService interface {
	ComputeWholeLifeCarbonSync(entities []calculator.CarbonCalculator) float64
	ComputeTotalCarbonConcurrent(entities []calculator.CarbonCalculator) float64
	ComputeEmbodiedCarbonSync(entities []calculator.EmbodiedCarbonCalculator) float64
}

type calculationService struct{}

func NewCalculationService() *calculationService {
	return &calculationService{}
}

func (s *calculationService) ComputeWholeLifeCarbonSync(entities []calculator.CarbonCalculator) float64 {
	var total float64
	for _, entity := range entities {
		total += entity.ComputeWholeLifeCarbon()
	}
	return total
}

func (s *calculationService) ComputeEmbodiedCarbonSync(entities ...calculator.EmbodiedCarbonCalculator) float64 {
	var total float64
	for _, entity := range entities {
		total += entity.CalculateEmbodiedCarbon()
	}
	return total
}

func (s *calculationService) ComputeTotalCarbonConcurrent(entities []calculator.CarbonCalculator) float64 {
	var wg sync.WaitGroup
	totalCarbon := make(chan float64)

	for _, entity := range entities {
		wg.Add(1)
		go func(entity calculator.CarbonCalculator) {
			defer wg.Done()
			totalCarbon <- entity.ComputeWholeLifeCarbon()
		}(entity)
	}

	go func() {
		wg.Wait()
		close(totalCarbon)
	}()

	var total float64
	for impact := range totalCarbon {
		total += impact
	}
	return total
}
