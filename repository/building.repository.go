package repository

import (
	"carbon-service/model"

	"gorm.io/gorm"
)

type BuildingRepository interface {
	Save(user *model.Building)
	ExistsByBuildingName(username string) bool
	FindByID(id int) (*model.Building, error)
	FindAll() ([]model.Building, error)
}

type buildingRepository struct {
	db *gorm.DB
}

func (repository *buildingRepository) Save(user *model.Building) {
	repository.db.Save(user)
}

// ExistsByUsername checks if a user with the given username exists in the database
func (repository *buildingRepository) ExistsByBuildingName(buildingName string) bool {
	var count int64
	repository.db.Model(&model.Building{}).Where("name = ?", buildingName).Count(&count)
	return count > 0
}

// Find by id returns a building by its id
func (repository *buildingRepository) FindByID(id int) (*model.Building, error) {
	var building model.Building
	err := repository.db.First(&building, id).Error
	if err != nil {
		return nil, err
	}
	return &building, nil
}

// FindAll returns all buildings
func (repository *buildingRepository) FindAll() ([]model.Building, error) {
	var buildings []model.Building
	err := repository.db.Find(&buildings).Error
	if err != nil {
		return nil, err
	}
	return buildings, nil
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *gorm.DB) BuildingRepository {

	repository := &buildingRepository{
		db: db,
	}

	return repository
}
