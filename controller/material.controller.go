package controller

import (
	"carbon-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MaterialController manages material-related HTTP handlers.
type MaterialController interface {
	createMaterial(ctx *gin.Context)
	getMaterial(ctx *gin.Context)
	getMaterials(ctx *gin.Context)
	getTotalCarbon(ctx *gin.Context)
}

type materialController struct {
	materialService    service.MaterialService
	calculationService service.CalculationService
}

// NewMaterialController sets up routes and handlers for material operations.
func NewMaterialController(router *gin.Engine, ms service.MaterialService, cs service.CalculationService) {
	mc := &materialController{
		materialService:    ms,
		calculationService: cs,
	}

	router.POST("/materials", mc.createMaterial)
	router.GET("/materials/:id", mc.getMaterial)
	router.GET("/materials", mc.getMaterials)
	router.GET("/materials/:id/total-carbon", mc.getTotalCarbon)
}

func (mc *materialController) createMaterial(ctx *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	material, err := mc.materialService.CreateMaterial(req.Name)
	if err != nil {
		respondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, material)
}

func (mc *materialController) getMaterial(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}
	material, err := mc.materialService.GetMaterial(uint(id))
	if err != nil {
		respondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, material)
}

func (mc *materialController) getMaterials(ctx *gin.Context) {
	materials, err := mc.materialService.GetAllMaterials()
	if err != nil {
		respondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, materials)
}

func (mc *materialController) getTotalCarbon(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}
	carbon, err := mc.materialService.ComputeTotalCarbon(uint(id))
	if err != nil {
		respondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"total_carbon": carbon})
}

func respondWithError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{"error": message})
}
