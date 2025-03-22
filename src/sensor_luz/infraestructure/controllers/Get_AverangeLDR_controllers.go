package controllers

import (
	"Multidiciplinario/src/sensor_luz/application/use_case"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAverageLightLDRController struct {
	useCase *use_case.GetAverageLightLDR
}

func NewGetAverageLightLDRController(useCase *use_case.GetAverageLightLDR) *GetAverageLightLDRController {
	return &GetAverageLightLDRController{useCase: useCase}
}

func (gt *GetAverageLightLDRController) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)

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
