package service

import (
	"carbon-service/model"
	"carbon-service/repository"
	"fmt"
)

// CalculationService interface for calculating carbon footprint of a building
type CalculationService interface {
	ComputeBuildingTotalCarbon(buildingID uint) (float64, error)
}

// CalculationService struct that implements the CalculationService interface
type calculationService struct {
	buildingRepo repository.BuildingRepository
}

// Creates a new CalculationService
func NewCalculationService(buildingRepo repository.BuildingRepository) *calculationService {
	return &calculationService{buildingRepo: buildingRepo}
}

// Calculate the total carbon footprint of a building
// This method fetches the building from the repository and calls the ComputeCarbonImpact() method on it, if the building is found successfully in the repository
// If the building is not found, an error is returned
func (s *calculationService) ComputeBuildingTotalCarbon(buildingID uint) (float64, error) {
	building, err := s.buildingRepo.FindByID(buildingID)
	if err != nil {
		return 0, fmt.Errorf(" Errorr in ComputeBuildingTotalCarbon, failed to find building with ID %d: %w", buildingID, err)
	}
	return building.ComputeCarbonImpact(), nil
}

// BuildingService interface for building operations like create, get, get all
type BuildingService interface {
	CreateBuilding(name string) (*model.Building, error)
	GetBuilding(id uint) (*model.Building, error)
	GetAllBuildings() ([]model.Building, error)
}

// BuildingService struct that implements the BuildingService interface
type buildingService struct {
	repo repository.BuildingRepository
}

// NewBuildingService creates a new BuildingService with the provided repository
func NewBuildingService(repo repository.BuildingRepository) BuildingService {
	return &buildingService{repo: repo}
}

// CreateBuilding creates a new building with the given name and returns it
// If a building with the same name already exists, an error is returned
// If the building is successfully created, it is returned
// If there is an error while saving the building, an error is returned
func (s *buildingService) CreateBuilding(name string) (*model.Building, error) {
	if s.repo.ExistsByBuildingName(name) {
		return nil, fmt.Errorf("building name '%s' already exists", name)
	}

	building := &model.Building{Name: name}
	if err := s.repo.Save(building); err != nil {
		return nil, fmt.Errorf("failed to create building: %w", err)
	}
	return building, nil
}

// GetBuilding fetches a building with the given ID and returns it
// If the building is not found, an error is returned
func (s *buildingService) GetBuilding(id uint) (*model.Building, error) {
	building, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find building with ID %d: %w", id, err)
	}
	return building, nil
}

// GetAllBuildings fetches all buildings and returns them
// If there is an error while fetching the buildings, an error is returned
func (s *buildingService) GetAllBuildings() ([]model.Building, error) {
	buildings, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find buildings: %w", err)
	}
	return buildings, nil
}
