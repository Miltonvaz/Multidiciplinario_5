package controllers

import (
	"Multidiciplinario/src/sensor_luz/application/use_case"
	"Multidiciplinario/src/sensor_luz/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetLightLDRByIDController struct {
	useCase *use_case.GetMeasurementByID
}

func NewGetLightLDRByIDController(useCase *use_case.GetMeasurementByID) *GetLightLDRByIDController {
	return &GetLightLDRByIDController{useCase: useCase}
}
func (gh *GetLightLDRByIDController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	measurement, err := gh.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la temperatura", "details": err.Error()})
		return
	}
	if (measurement == entities.LightSensorLDR{}) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontró el registro con el ID proporcionado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": measurement,
	})
}
