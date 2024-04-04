package controller

import (
	"carbon-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BuildingController interface {
	CreateBuilding(ctx *gin.Context)
	GetBuilding(ctx *gin.Context)
	GetBuildings(ctx *gin.Context)
	GetTotalCarbon(ctx *gin.Context)
}

type buildingController struct {
	buildingService    service.BuildingService
	calculationService service.CalculationService
}

func NewBuildingController(router *gin.Engine, bs service.BuildingService, cs service.CalculationService) {
	controller := &buildingController{
		buildingService:    bs,
		calculationService: cs,
	}

	router.POST("/buildings", controller.CreateBuilding)
	router.GET("/buildings/:id", controller.GetBuilding)
	router.GET("/buildings", controller.GetBuildings)
	// Assuming ComputeTotalCarbon is part of the BuildingService for simplicity.
	router.GET("/buildings/:id/total-carbon", controller.GetTotalCarbon)
}

func (c *buildingController) CreateBuilding(ctx *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	building, err := c.buildingService.CreateBuilding(req.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, building)
}

func (c *buildingController) GetBuilding(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}
	building, err := c.buildingService.GetBuilding(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Building not found", "id": id})
		return
	}
	ctx.JSON(http.StatusOK, building)
}

func (bc *buildingController) GetBuildings(ctx *gin.Context) {
	buildings, err := bc.buildingService.GetAllBuildings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching buildings"})
		return
	}
	ctx.JSON(http.StatusOK, buildings)
}

func (bc *buildingController) GetTotalCarbon(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	totalCarbon, err := bc.calculationService.ComputeBuildingTotalCarbon(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Building not found or calculation error", "id": id})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"totalCarbon": totalCarbon})
}
