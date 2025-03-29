package controllers

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/application/use_case"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetTemperatureByIDController struct {
	useCase *use_case.GetMeasurementByID
}

func NewGetTemperatureByIDController(useCase *use_case.GetMeasurementByID) *GetTemperatureByIDController {
	return &GetTemperatureByIDController{useCase: useCase}
}
func (gh *GetTemperatureByIDController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userIDStr := ctx.Param("userID")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido"})
		return
	}

	measurement, err := gh.useCase.Execute(id, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la temperatura", "details": err.Error()})
		return
	}
	if (measurement == entities.BodyTemperature{}) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontró el registro con el ID proporcionado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": measurement,
	})
}
