package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"carbon-service/controller"
	"carbon-service/database"
	"carbon-service/helpers"
	"carbon-service/model"
	"carbon-service/repository"
	"carbon-service/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
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

	// Connect to the database
	db, err := database.ConnectDatabase(&dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Perform database migration
	if err := db.AutoMigrate(&model.Building{}, &model.Assembly{}, &model.Material{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schemas: %v", err)
	}

	// Initialize repository, service, and controller
	br := repository.NewBuildingRepository(db)
	ar := repository.NewAssemblyRepository(db)
	mr := repository.NewMaterialRepository(db)

	// Initialize services
	cs := service.NewCalculationService() // Assuming this is correctly implemented based on previous discussions

	// Inject dependencies into building service
	bs := service.NewBuildingService(br, cs)
	as := service.NewAssemblyService(ar, cs)
	ms := service.NewMaterialService(mr, cs)

	// Initialize the router which will handle the requests
	router := gin.Default()

	// Inject services into controller
	controller.NewBuildingController(router, bs, cs)
	controller.NewAssemblyController(router, as, cs)
	controller.NewMaterialController(router, ms, cs)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
