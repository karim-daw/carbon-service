package service

import (
	"carbon-service/repository"
	"errors"
)

type CarbonCalculationService interface {
	// ComputeTotalCarbon calculates the total carbon footprint of a building by its ID.
	ComputeTotalCarbon(id int) (float64, error)
}

// carbonCalculationService is a concrete implementation of CarbonCalculationService.
type carbonCalculationService struct {
	buildingRepo repository.BuildingRepository
}

// NewCarbonCalculationService creates a new instance of a service capable of calculating carbon metrics.
func NewCarbonCalculationService(repo repository.BuildingRepository) CarbonCalculationService {
	return &carbonCalculationService{
		buildingRepo: repo,
	}
}

// ComputeTotalCarbon calculates the total carbon footprint of a building by its ID.
func (service *carbonCalculationService) ComputeTotalCarbon(id int) (float64, error) {
	building, err := service.buildingRepo.FindByID(id)
	if err != nil {
		return 0, errors.New("building not found")
	}
	return building.EC + building.OC, nil
}
