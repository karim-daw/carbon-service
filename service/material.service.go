package service

import (
	"carbon-service/model"
	"carbon-service/repository"
	"fmt"
)

// MaterialService defines the operations available for managing materials,
// including creation, retrieval, and carbon footprint calculation.
type MaterialService interface {
	CreateMaterial(name string) (*model.Material, error)
	GetMaterial(id uint) (*model.Material, error)
	GetAllMaterials() ([]model.Material, error)
	ComputeTotalCarbon(materialID uint) (float64, error)
}

// materialService provides a concrete implementation of the MaterialService,
// interacting with material data and carbon calculations.
type materialService struct {
	repo              repository.MaterialRepository
	carbonCalcService CalculationService // Dependency for carbon calculations
}

// ComputeTotalCarbon implements MaterialService.
func (m *materialService) ComputeTotalCarbon(materialID uint) (float64, error) {
	var material *model.Material
	material, err := m.repo.FindByID(materialID)
	if err != nil {
		return 0, fmt.Errorf("failed to find material with ID %d: %w", materialID, err)
	}
	return material.ComputeWholeLifeCarbon(), nil
}

// CreateMaterial implements MaterialService.
func (m *materialService) CreateMaterial(name string) (*model.Material, error) {
	if m.repo.ExistsByMaterialName(name) {
		return nil, fmt.Errorf("material name '%s' already exists", name)
	}

	material := &model.Material{Name: name}
	if err := m.repo.Save(material); err != nil {
		return nil, fmt.Errorf("failed to create material: %w", err)
	}
	return material, nil
}

// GetAllMaterials implements MaterialService.
func (m *materialService) GetAllMaterials() ([]model.Material, error) {
	materials, err := m.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find materials: %w", err)
	}
	return materials, nil
}

// GetMaterial implements MaterialService.
func (m *materialService) GetMaterial(id uint) (*model.Material, error) {
	material, err := m.repo.EagerFindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find material with ID %d: %w", id, err)
	}
	return material, nil
}

// NewMaterialService initializes a new material service with necessary dependencies.
func NewMaterialService(r repository.MaterialRepository, cs CalculationService) MaterialService {
	return &materialService{
		repo:              r,
		carbonCalcService: cs,
	}
}
