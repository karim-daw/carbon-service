package controller

import (
	"carbon-service/service"
	"net/http"
	"strconv"

	"carbon-service/helpers"

	"github.com/gin-gonic/gin"
)

type buildingController struct {
	buildingService    service.BuildingService
	calculationService service.CalculationService
}

// NewBuildingController sets up routes and handlers for building operations.
func NewBuildingController(router *gin.Engine, bs service.BuildingService, cs service.CalculationService) {
	bc := &buildingController{
		buildingService:    bs,
		calculationService: cs,
	}

	router.POST("/buildings", bc.createBuilding)
	router.PUT("/buildings/:id", bc.updateBuilding)
	router.GET("/buildings/:id", bc.getBuilding)
	router.GET("/buildings", bc.getBuildings)
	router.GET("/buildings/:id/calculation/total-carbon", bc.getTotalCarbon)
	router.GET("/buildings/:id/calculation/embodied-carbon", bc.getEmbodiedCarbon)
}

// createBuilding handles the creation of a new building with the provided data.
// It expects a JSON payload with a name and an array of assemblies.
// endpoint: POST /buildings
func (bc *buildingController) createBuilding(ctx *gin.Context) {
	var req service.CreateBuildingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	building, err := bc.buildingService.CreateBuilding(req)
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, building)
}

// getBuilding fetches a building by its ID.
// endpoint: GET /buildings/:id
func (bc *buildingController) getBuilding(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}
	building, err := bc.buildingService.GetBuilding(uint(id))
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusNotFound, "Building not found")
		return
	}
	ctx.JSON(http.StatusOK, building)
}

// getBuildings fetches all buildings stored in the system.
// endpoint: GET /buildings
func (bc *buildingController) getBuildings(ctx *gin.Context) {
	buildings, err := bc.buildingService.GetAllBuildings()
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusInternalServerError, "Error fetching buildings")
		return
	}
	ctx.JSON(http.StatusOK, buildings)
}

// getTotalCarbon fetches the total carbon impact of a building by its ID.
// endpoint: GET /buildings/:id/calculation/total-carbon
func (bc *buildingController) getTotalCarbon(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}
	totalCarbon, err := bc.buildingService.ComputeTotalCarbon(uint(id))
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusNotFound, "Building not found or calculation error")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"totalCarbon": totalCarbon})
}

// getEmbodiedCarbon fetches the embodied carbon of a building by its ID.
// endpoint: GET /buildings/:id/calculation/embodied-carbon
func (bc *buildingController) getEmbodiedCarbon(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}
	embodiedCarbon, err := bc.buildingService.ComputeEmbodiedCarbon(uint(id))
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusNotFound, "Building not found or calculation error")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"embodiedCarbon": embodiedCarbon})
}

// put endpoint: PUT /buildings/:id
func (bc *buildingController) updateBuilding(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}
	var req service.UpdateBuildingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}
	building, err := bc.buildingService.UpdateBuilding(uint(id), req)
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, building)
}
