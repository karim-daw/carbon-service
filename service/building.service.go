package service

import (
	"carbon-service/model"
	"carbon-service/repository"
	"errors"
)

type BuildingService interface {
	CreateBuilding(name string, location string, ec float64, oc float64) error
	GetBuilding(id int) (*model.Building, error)
	GetBuildings() ([]model.Building, error)
	ComputeTotalCarbon(id int) float64
}

type buildingService struct {
	repository repository.BuildingRepository
}

func (service *buildingService) CreateBuilding(name string, location string, ec float64, oc float64) error {

	exists := service.repository.ExistsByBuildingName(name)
	if exists {
		return errors.New("building name already exists")
	}
	service.repository.Save(&model.Building{
		Name:     name,
		Location: location,
		EC:       ec,
		OC:       oc,
	})
	return nil
}

// GetBuilding returns a building by its id
func (service *buildingService) GetBuilding(id int) (*model.Building, error) {
	building, err := service.repository.FindByID(id)
	if err != nil {
		return nil, errors.New("building not found")
	}
	return building, nil
}

// GetBuildings returns all buildings
func (service *buildingService) GetBuildings() ([]model.Building, error) {
	buildings, err := service.repository.FindAll()
	if err != nil {
		return nil, errors.New("error fetching buildings from repository")
	}
	return buildings, nil
}

// ComputeTotalCarbon returns the sum of embodied carbon and operational carbon for a building
func (service *buildingService) ComputeTotalCarbon(id int) float64 {
	building, err := service.GetBuilding(id)
	if err != nil {
		return 0
	}
	return building.EC + building.OC
}

// NewBuildingService creates a new BuildingService
func NewBuildingService(repository repository.BuildingRepository) BuildingService {
	return &buildingService{
		repository: repository,
	}
}
