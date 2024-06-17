// service/calculation.go

package service

import (
	"carbon-service/model"
	"sync"
)

type CalculationService interface {
	ComputeWholeLifeCarbonSync(entities ...model.CarbonCalculator) float64
	ComputeTotalCarbonConcurrent(entities ...model.CarbonCalculator) float64
	ComputeEmbodiedCarbonSync(entities ...model.EmbodiedCarbonCalculator) float64
}

type calculationService struct{}

func NewCalculationService() *calculationService {
	return &calculationService{}
}

func (s *calculationService) ComputeWholeLifeCarbonSync(entities ...model.CarbonCalculator) float64 {
	var total float64
	for _, entity := range entities {
		total += entity.ComputeWholeLifeCarbon()
	}
	return total
}

func (s *calculationService) ComputeEmbodiedCarbonSync(entities ...model.EmbodiedCarbonCalculator) float64 {
	var total float64
	for _, entity := range entities {
		total += entity.CalculateEmbodiedCarbon()
	}
	return total
}

func (s *calculationService) ComputeTotalCarbonConcurrent(entities ...model.CarbonCalculator) float64 {
	var wg sync.WaitGroup
	totalCarbon := make(chan float64)

	for _, entity := range entities {
		wg.Add(1)
		go func(entity model.CarbonCalculator) {
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
