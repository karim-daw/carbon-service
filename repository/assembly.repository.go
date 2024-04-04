package repository

import (
	"carbon-service/model"

	"gorm.io/gorm"
)

// AssemblyRepository is an interface for interacting with the assemblies table.
type AssemblyRepository interface {
	Save(assembly *model.Assembly) error
	FindByID(id uint) (*model.Assembly, error)
	EagerFindByID(id uint) (*model.Assembly, error)
	FindAll() ([]model.Assembly, error)
	EagerFindAll() ([]model.Assembly, error)
	ExistsByAssemblyName(name string) bool
}

// assemblyRepository is a concrete implementation of AssemblyRepository.
type assemblyRepository struct {
	db *gorm.DB
}

// Save persists an assembly to the database.
func (r *assemblyRepository) Save(assembly *model.Assembly) error {
	return r.db.Save(assembly).Error
}

// FindByID retrieves an assembly from the database based on the provided ID.
// It returns a pointer to the found assembly and an error, if any.
func (r *assemblyRepository) FindByID(id uint) (*model.Assembly, error) {
	var assembly model.Assembly
	err := r.db.First(&assembly, id).Error
	if err != nil {
		return nil, err
	}
	return &assembly, nil
}

// EagerFindByID retrieves an assembly from the database based on the provided ID,
// preloading its materials.
// It returns a pointer to the found assembly and an error, if any.
func (r *assemblyRepository) EagerFindByID(id uint) (*model.Assembly, error) {
	var assembly model.Assembly
	// pre load materials and all buildings that use this assembly
	err := r.db.Preload("Materials").Preload("Buildings").First(&assembly, id).Error
	// err := r.db.Preload("Materials").First(&assembly, id).Error
	if err != nil {
		return nil, err
	}
	return &assembly, nil
}

// FindAll retrieves all assemblies from the database.
// It returns a slice of assemblies and an error, if any.
func (r *assemblyRepository) FindAll() ([]model.Assembly, error) {
	var assemblies []model.Assembly
	err := r.db.Find(&assemblies).Error
	if err != nil {
		return nil, err
	}
	return assemblies, nil
}

// EagerFindAll retrieves all assemblies from the database, preloading their materials.
// It returns a slice of assemblies and an error, if any.
func (r *assemblyRepository) EagerFindAll() ([]model.Assembly, error) {
	var assemblies []model.Assembly
	err := r.db.Preload("Materials").Find(&assemblies).Error
	if err != nil {
		return nil, err
	}
	return assemblies, nil
}

// ExistsByAssemblyName checks if an assembly with the provided name exists in the database.
// It returns true if the assembly exists, false otherwise.
func (r *assemblyRepository) ExistsByAssemblyName(name string) bool {
	var assembly model.Assembly
	err := r.db.Where("name = ?", name).First(&assembly).Error
	return err == nil
}

// NewAssemblyRepository creates a new AssemblyRepository with the provided database connection.
func NewAssemblyRepository(db *gorm.DB) AssemblyRepository {
	return &assemblyRepository{db}
}
