package controllers

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAverageHeartRateController struct {
	useCase *use_case.GetAverageHeartRate
}

func NewGetAverageHeartRateController(useCase *use_case.GetAverageHeartRate) *GetAverageHeartRateController {
	return &GetAverageHeartRateController{useCase: useCase}
}

func (gt *GetAverageHeartRateController) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID inválido"})
		return
	}

	averageHeartRate, err := gt.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al calcular el promedio de ritmo cardíaco",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"average_heart_rate": averageHeartRate})
}
