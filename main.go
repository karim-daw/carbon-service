package main

import (
	"carbon-service/model"
	"log"

	"carbon-service/controller"
	"carbon-service/database"
	"carbon-service/helpers"
	"carbon-service/repository"
	"carbon-service/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// some go cli commands
// go mod init carbon-service
// go mod tidy
// go mod vendor
// go mod download
// go mod verify
// go mod graph
// go mod edit
// go mod why
// go mod why -m github.com/gin-gonic/gin
// go mod why -m github.com/jinzhu/gorm
// run the app
// go run main.go
func main() {
	godotenv.Load()

	// create db config
	dbConfig := &database.DbConfig{
		User:     helpers.LoadEnvVar("POSTGRES_USERNAME"),
		Password: helpers.LoadEnvVar("POSTGRES_PASSWORD"),
		DbName:   helpers.LoadEnvVar("DATABASE_NAME"),
		Host:     helpers.LoadEnvVar("DATABASE_HOST"),
		Port:     helpers.LoadEnvVar("DATABASE_PORT"),
		Schema:   helpers.LoadEnvVar("DATABASE_SCHEMA"),
	}

	db, err := database.ConnectDatabase(dbConfig)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Building{})
	if err != nil {
		panic(err)
	}

	buildingRepository := repository.NewBuildingRepository(db)
	buildingService := service.NewBuildingService(buildingRepository)
	calculationService := service.NewCarbonCalculationService(buildingRepository)

	engine := gin.Default()

	controller.NewBuildingController(engine, buildingService, calculationService)

	log.Fatal(engine.Run(":8080"))
}
