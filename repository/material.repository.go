package repository

import (
	"carbon-service/model"

	"gorm.io/gorm"
)

// MaterialRepository is an interface for interacting with the materials table.
type MaterialRepository interface {
	Save(material *model.Material) error
	ExistsByMaterialName(materialName string) bool
	FindByID(id uint) (*model.Material, error)
	EagerFindByID(id uint) (*model.Material, error)
	FindAll() ([]model.Material, error)
}

// materialRepository is a concrete implementation of MaterialRepository.
type materialRepository struct {
	db *gorm.DB
}

// Save persists a material to the database.
func (r *materialRepository) Save(material *model.Material) error {
	return r.db.Save(material).Error
}

// ExistsByMaterialName checks if a material with the provided name exists in the database.
func (r *materialRepository) ExistsByMaterialName(materialName string) bool {
	var count int64
	r.db.Model(&model.Material{}).Where("name = ?", materialName).Count(&count)
	return count > 0
}

// FindByID retrieves a material from the database based on the provided ID.
// It returns a pointer to the found material and an error, if any.
func (r *materialRepository) FindByID(id uint) (*model.Material, error) {
	var material model.Material
	err := r.db.First(&material, id).Error
	if err != nil {
		return nil, err
	}
	return &material, nil
}

// EagerFindByID retrieves a material from the database based on the provided ID,
// preloading its assemblies.
// It returns a pointer to the found material and an error, if any.
func (r *materialRepository) EagerFindByID(id uint) (*model.Material, error) {
	var material model.Material
	err := r.db.Preload("Assemblies").First(&material, id).Error
	if err != nil {
		return nil, err
	}
	return &material, nil
}

// FindAll retrieves all materials from the database.
// It returns a slice of materials and an error, if any.
func (r *materialRepository) FindAll() ([]model.Material, error) {
	var materials []model.Material
	err := r.db.Find(&materials).Error
	if err != nil {
		return nil, err
	}
	return materials, nil
}

// NewMaterialRepository creates a new MaterialRepository with the provided database connection.
func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepository{db}
}
