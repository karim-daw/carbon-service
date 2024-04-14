// pkg/repository/building_repository_test.go

package tests

import (
	"carbon-service/database"
	"carbon-service/helpers"
	"carbon-service/model"
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseTestSuite is the test suite.
type DatabaseTestSuite struct {
	suite.Suite
	db *gorm.DB
}

// SetupSuite is called once before the test suite runs.
func (suite *DatabaseTestSuite) SetupSuite() {

	// Load environment variables
	if err := godotenv.Load("/Users/karimdaw/Projects/carbon-service/.env"); err != nil {
		log.Println("Warning: .env file could not be loaded. Proceeding with environment variables.")
	}
	// Create db config from environment variables
	dbConfig := database.DbConfig{
		User:     helpers.LoadEnvVar("POSTGRES_USERNAME"),
		Password: helpers.LoadEnvVar("POSTGRES_PASSWORD"),
		DbName:   helpers.LoadEnvVar("DATABASE_NAME"),
		Host:     helpers.LoadEnvVar("DATABASE_HOST"),
		Port:     helpers.LoadEnvVar("DATABASE_PORT"),
		Schema:   helpers.LoadEnvVar("DATABASE_SCHEMA"),
	}

	// Set up a PostgreSQL database for testing
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, "carbon-db-test", dbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	suite.Require().NoError(err, "Error connecting to the test database")

	// Enable logging for Gorm during tests
	suite.db = db.Debug()

	if err := db.AutoMigrate(&model.Building{}, &model.Assembly{}, &model.Material{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schemas: %v", err)
	}

	// Auto-migrate tables
	suite.Require().NoError(err, "Error auto-migrating database tables")
}

// TestUserInsertion tests inserting a user record.
func (suite *DatabaseTestSuite) TestBuildingEntity() {

	// create new building with all required fields
	building := model.Building{
		Name:                  "TestBuilding 1",
		FTF:                   4.0,
		GroundFloorArea:       1000.0,
		WWR:                   0.4,
		AboveGroundFloorCount: 2,
		UnderGroundFloorCount: 1,
	}

	err := suite.db.Create(&building).Error
	suite.Require().NoError(err, "Error creating user record")

	// Retrieve the inserted user
	var retrievedBuilding model.Building
	err = suite.db.First(&retrievedBuilding, "name = ?", "TestBuilding 1").Error
	suite.Require().NoError(err, "Error retrieving user record")

	// Verify that the retrieved user matches the inserted user
	suite.Equal(building.Name, retrievedBuilding.Name, "Names should match")
}

// TestAssemblyEntity tests inserting an assembly record.
func (suite *DatabaseTestSuite) TestAssemblyEntity() {

	// create new assembly with all required fields
	assembly := model.Assembly{
		Name: "TestAssembly 1",
	}

	err := suite.db.Create(&assembly).Error
	suite.Require().NoError(err, "Error creating assembly record")

	// Retrieve the inserted assembly
	var retrievedAssembly model.Assembly
	err = suite.db.First(&retrievedAssembly, "name = ?", "TestAssembly 1").Error
	suite.Require().NoError(err, "Error retrieving assembly record")

	// Verify that the retrieved assembly matches the inserted assembly
	suite.Equal(assembly.Name, retrievedAssembly.Name, "Names should match")
}

// TestMaterialEntity tests inserting a material record.
func (suite *DatabaseTestSuite) TestMaterialEntity() {

	// create new material with all required fields
	material := model.Material{
		Name: "TestMaterial 1",
	}

	err := suite.db.Create(&material).Error
	suite.Require().NoError(err, "Error creating material record")

	// Retrieve the inserted material
	var retrievedMaterial model.Material
	err = suite.db.First(&retrievedMaterial, "name = ?", "TestMaterial 1").Error
	suite.Require().NoError(err, "Error retrieving material record")

	// Verify that the retrieved material matches the inserted material
	suite.Equal(material.Name, retrievedMaterial.Name, "Names should match")
}

// TearDownSuite is called once after the test suite runs.
func (suite *DatabaseTestSuite) TearDownSuite() {
	// Clean up: Close the database connection

	// Drop the tables
	// err := suite.db.Migrator().DropTable(&model.Building{}, &model.Assembly{}, &model.Material{})
	// suite.Require().NoError(err, "Error dropping database tables")

	sqlDB, err := suite.db.DB()
	sqlDB.Close()

	suite.Require().NoError(err, "Error closing the test database")
	log.Println("Test suite teardown completed successfully.")

}

// TestSuite runs the test suite.
func TestSuite(t *testing.T) {
	// Skip the tests if the PostgreSQL connection details are not provided

	suite.Run(t, new(DatabaseTestSuite))

	// teardown

}
