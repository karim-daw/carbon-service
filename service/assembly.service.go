package service

import (
	"carbon-service/model"
	"carbon-service/repository"
	"fmt"
)

// AssemblyService defines the operations available for managing assemblies,
// including creation, retrieval, and material addition.
type AssemblyService interface {
	CreateAssembly(name string) (*model.Assembly, error)
	GetAssembly(id uint) (*model.Assembly, error)
	GetAllAssemblies() ([]model.Assembly, error)
	ComputeTotalCarbon(assemblyID uint) (float64, error)
}

// assemblyService provides a concrete implementation of the AssemblyService,
// interacting with assembly data and carbon calculations.
type assemblyService struct {
	repo              repository.AssemblyRepository
	carbonCalcService CalculationService // Dependency for carbon calculations
}

// ComputeTotalCarbon implements AssemblyService.
func (as *assemblyService) ComputeTotalCarbon(assemblyID uint) (float64, error) {
	var assembly *model.Assembly
	assembly, err := as.repo.EagerFindByID(assemblyID)
	if err != nil {
		return 0, fmt.Errorf("failed to find assembly with ID %d: %w", assemblyID, err)
	}
	return assembly.ComputeWholeLifeCarbon(), nil
}

// CreateAssembly implements AssemblyService.
func (as *assemblyService) CreateAssembly(name string) (*model.Assembly, error) {
	if as.repo.ExistsByAssemblyName(name) {
		return nil, fmt.Errorf("assembly name '%s' already exists", name)
	}

	assembly := &model.Assembly{Name: name}
	if err := as.repo.Save(assembly); err != nil {
		return nil, fmt.Errorf("failed to create assembly: %w", err)
	}
	return assembly, nil

}

// GetAllAssemblies implements AssemblyService.
func (a *assemblyService) GetAllAssemblies() ([]model.Assembly, error) {
	assemblies, err := a.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find assemblies: %w", err)
	}
	return assemblies, nil
}

// GetAssembly implements AssemblyService.
func (a *assemblyService) GetAssembly(id uint) (*model.Assembly, error) {
	assembly, err := a.repo.EagerFindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find assembly with ID %d: %w", id, err)
	}
	return assembly, nil
}

// NewAssemblyService initializes a new assembly service with necessary dependencies.
func NewAssemblyService(r repository.AssemblyRepository, cs CalculationService) AssemblyService {
	return &assemblyService{
		repo:              r,
		carbonCalcService: cs,
	}
}
