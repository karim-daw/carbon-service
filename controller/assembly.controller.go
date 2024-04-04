package controller

import (
	"carbon-service/service"
	"net/http"
	"strconv"

	"carbon-service/helpers"

	"github.com/gin-gonic/gin"
)

type assemblyController struct {
	assemblyService    service.AssemblyService
	calculationService service.CalculationService
}

// NewAssemblyController sets up routes and handlers for assembly operations.
func NewAssemblyController(router *gin.Engine, as service.AssemblyService, cs service.CalculationService) {
	ac := &assemblyController{
		assemblyService:    as,
		calculationService: cs,
	}

	router.POST("/assemblies", ac.createAssembly)
	router.GET("/assemblies/:id", ac.getAssembly)
	router.GET("/assemblies", ac.getAssemblies)
	router.GET("/assemblies/:id/total-carbon", ac.getTotalCarbon)
}

func (ac *assemblyController) createAssembly(ctx *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	assembly, err := ac.assemblyService.CreateAssembly(req.Name)
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, assembly)
}

func (ac *assemblyController) getAssembly(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}
	assembly, err := ac.assemblyService.GetAssembly(uint(id))
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusNotFound, "Assembly not found")
		return
	}
	ctx.JSON(http.StatusOK, assembly)
}

func (ac *assemblyController) getAssemblies(ctx *gin.Context) {
	assemblies, err := ac.assemblyService.GetAllAssemblies()
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, assemblies)
}

func (ac *assemblyController) getTotalCarbon(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}
	carbon, err := ac.assemblyService.ComputeTotalCarbon(uint(id))
	if err != nil {
		helpers.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"total_carbon": carbon})
}
