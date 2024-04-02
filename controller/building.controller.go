package controller

import (
	"carbon-service/model"
	"carbon-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BuildingController interface {
	CreateBuilding(context *gin.Context)
	GetBuilding(context *gin.Context)
	GetBuildings(context *gin.Context)
	GetTotalCarbon(context *gin.Context)
}

type buildingController struct {
	service                  service.BuildingService
	carbonCalculationService service.CarbonCalculationService
}

// CreateBuilding adds a building from JSON received in the request body.
// sample request: POST /buildings
func (controller buildingController) CreateBuilding(context *gin.Context) {
	var building model.Building
	if err := context.ShouldBindJSON(&building); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := controller.service.CreateBuilding(building.Name, building.Location, building.EC, building.OC)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	context.JSON(201, nil)
}

// GetBuilding responds with a building by ID.
// sample request: GET /buildings/1
func (controller buildingController) GetBuilding(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	building, err := controller.service.GetBuilding(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Building not found", "id": id})
		return
	}
	context.JSON(http.StatusOK, building)
}

// GetBuildings responds with the list of all buildings as JSON.
// sample request: GET /buildings
func (controller buildingController) GetBuildings(context *gin.Context) {
	// get list of buildings
	buildings, err := controller.service.GetBuildings()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching buildings"})
		return
	}
	context.JSON(http.StatusOK, buildings)
}

// GetTotalCarbon responds with the sum of embodied carbon and operational carbon for a building.
// sample request: GET /buildings/1/total-carbon
func (controller buildingController) GetTotalCarbon(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	totalCarbon, err := controller.carbonCalculationService.ComputeTotalCarbon(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Building not found", "id": id})
		return
	}
	context.JSON(http.StatusOK, gin.H{"totalCarbon": totalCarbon})
}

func NewBuildingController(engine *gin.Engine,
	buildingService service.BuildingService,
	carbonCalculationService service.CarbonCalculationService,
) {
	controller := &buildingController{
		service:                  buildingService,
		carbonCalculationService: carbonCalculationService,
	}

	engine.POST("/buildings", controller.CreateBuilding)
	engine.GET("/buildings/:id", controller.GetBuilding)
	engine.GET("/buildings", controller.GetBuildings)
	engine.GET("/buildings/:id/total-carbon", controller.GetTotalCarbon)

}
