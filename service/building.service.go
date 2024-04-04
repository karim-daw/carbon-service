package service

import (
	"carbon-service/model"
	"carbon-service/repository"
	"fmt"
)

// BuildingService defines the operations available for managing buildings,
// including creation, retrieval, and carbon footprint calculation.
type BuildingService interface {
	CreateBuilding(req CreateBuildingRequest) (*model.Building, error)
	GetBuilding(id uint) (*model.Building, error)
	GetAllBuildings() ([]model.Building, error)
	ComputeTotalCarbon(buildingID uint) (float64, error)
}

// buildingService provides a concrete implementation of the BuildingService,
// interacting with building data and carbon calculations.
type buildingService struct {
	repo              repository.BuildingRepository
	carbonCalcService CalculationService // Dependency for carbon calculations
}

// NewBuildingService initializes a new building service with necessary dependencies.
func NewBuildingService(r repository.BuildingRepository, cs CalculationService) BuildingService {
	return &buildingService{
		repo:              r,
		carbonCalcService: cs,
	}
}

type CreateBuildingRequest struct {
	Name       string           `json:"name" binding:"required"`
	Assemblies []model.Assembly `json:"assemblies"`
}

// CreateBuilding attempts to add a new building with the given name,
// ensuring name uniqueness within the repository.
func (bs *buildingService) CreateBuilding(req CreateBuildingRequest) (*model.Building, error) {
	if bs.repo.ExistsByBuildingName(req.Name) {
		return nil, fmt.Errorf("building name '%s' already exists", req.Name)
	}

	assemblies := make([]*model.Assembly, len(req.Assemblies))
	for i, assembly := range req.Assemblies {
		assemblies[i] = &assembly
	}

	building := &model.Building{
		Name:       req.Name,
		Assemblies: assemblies, // This can be an empty slice if no assemblies are provided
	}
	if err := bs.repo.Save(building); err != nil {
		return nil, fmt.Errorf("failed to create building: %w", err)
	}
	return building, nil
}

// GetBuilding fetches a single building by its unique identifier.
// this method uses the EagerFindByID method from the repository to preload
// the building's assemblies and materials.
func (bs *buildingService) GetBuilding(id uint) (*model.Building, error) {
	building, err := bs.repo.EagerFindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find building with ID %d: %w", id, err)
	}
	return building, nil
}

// GetAllBuildings retrieves all buildings stored in the repository.
func (bs *buildingService) GetAllBuildings() ([]model.Building, error) {
	buildings, err := bs.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find buildings: %w", err)
	}
	return buildings, nil
}

// Example of a method in the buildingService that preloads necessary data before calculation
func (bs *buildingService) ComputeTotalCarbon(buildingID uint) (float64, error) {
	var building *model.Building
	// Preload Assemblies and Materials for the building
	building, err := bs.repo.EagerFindByID(buildingID) // Assign the value to building pointer
	if err != nil {
		return 0, fmt.Errorf("failed to find building with ID %d: %w", buildingID, err)
	}

	// Now that we have a fully loaded building, calculate the total carbon impact
	return building.ComputeWholeLifeCarbon(), nil
}
