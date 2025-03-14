package controllers

import (
	"Multidiciplinario/src/sensor_temperatura_humedad/application/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAverageHumidityController struct {
	useCase *use_case.GetAverageHumidity
}

func NewGetAverageHumidityController(useCase *use_case.GetAverageHumidity) *GetAverageHumidityController {
	return &GetAverageHumidityController{useCase: useCase}
}

func (gh *GetAverageHumidityController) Execute(ctx *gin.Context) {
	averageHumidity, err := gh.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al calcular el promedio de humedad",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"average_humidity": averageHumidity})
}
