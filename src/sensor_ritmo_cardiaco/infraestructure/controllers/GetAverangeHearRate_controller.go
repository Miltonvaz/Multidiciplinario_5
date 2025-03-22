package controllers

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAverageHeartRateController struct {
	useCase *use_case.GetAverageHeartRate
}

func NewGetAverageHeartRateController(useCase *use_case.GetAverageHeartRate) *GetAverageHeartRateController {
	return &GetAverageHeartRateController{useCase: useCase}
}

func (gt *GetAverageHeartRateController) Execute(ctx *gin.Context) {
	averageTemperature, err := gt.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al calcular el promedio de temperatura",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"average_temperature": averageTemperature})
}
