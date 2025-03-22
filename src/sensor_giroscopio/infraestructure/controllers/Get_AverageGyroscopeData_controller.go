package controllers

import (
	"Multidiciplinario/src/sensor_giroscopio/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAverageGyroscopeSensorController struct {
	useCase *use_case.GetAverageGyroscopeData
}

func NewGetAverageGyroscopeSensorController(useCase *use_case.GetAverageGyroscopeData) *GetAverageGyroscopeSensorController {
	return &GetAverageGyroscopeSensorController{useCase: useCase}
}

func (gt *GetAverageGyroscopeSensorController) Execute(ctx *gin.Context) {
	averageGyroscope, err := gt.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al calcular el promedio de giroscopio",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"average_gyroscope": averageGyroscope})
}
