package controllers

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/application/use_case"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAverageTemperatureController struct {
	useCase *use_case.GetAverageTemperature
}

func NewGetAverageTemperatureController(useCase *use_case.GetAverageTemperature) *GetAverageTemperatureController {
	return &GetAverageTemperatureController{useCase: useCase}
}

func (gt *GetAverageTemperatureController) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido"})
		return
	}

	averageTemperature, err := gt.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al calcular el promedio de temperatura",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"average_temperature": averageTemperature})
}
