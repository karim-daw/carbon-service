package repository

import (
	"carbon-service/model"

	"gorm.io/gorm"
)

type BuildingRepository interface {
	Save(building *model.Building) error
	ExistsByBuildingName(buildingName string) bool
	FindByID(id uint) (*model.Building, error)
	EagerFindByID(id uint) (*model.Building, error)
	FindAll() ([]model.Building, error)
	EagerFindAll() ([]model.Building, error)
}

type buildingRepository struct {
	db *gorm.DB
}

func (r *buildingRepository) Save(building *model.Building) error {
	return r.db.Save(building).Error
}

func (r *buildingRepository) ExistsByBuildingName(buildingName string) bool {
	var count int64
	r.db.Model(&model.Building{}).Where("name = ?", buildingName).Count(&count)
	return count > 0
}

// FindByID fetches a building by ID.
func (r *buildingRepository) FindByID(id uint) (*model.Building, error) {
	var building model.Building
	err := r.db.First(&building, id).Error
	if err != nil {
		return nil, err
	}
	return &building, nil
}

// EagerFindByID fetches a building by ID, preloading its assemblies and materials.
func (r *buildingRepository) EagerFindByID(id uint) (*model.Building, error) {
	var building model.Building
	err := r.db.Preload("Assemblies.Materials").First(&building, id).Error
	if err != nil {
		return nil, err
	}
	return &building, nil
}

func (r *buildingRepository) FindAll() ([]model.Building, error) {
	var buildings []model.Building
	err := r.db.Find(&buildings).Error
	if err != nil {
		return nil, err
	}
	return buildings, nil
}

func (r *buildingRepository) EagerFindAll() ([]model.Building, error) {
	var buildings []model.Building
	err := r.db.Preload("Assemblies.Materials").Find(&buildings).Error
	if err != nil {
		return nil, err
	}
	return buildings, nil
}

func NewBuildingRepository(db *gorm.DB) BuildingRepository {
	return &buildingRepository{db: db}
}
