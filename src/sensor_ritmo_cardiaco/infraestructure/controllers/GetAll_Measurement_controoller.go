package controllers

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAllHeartRateController struct {
	useCase *use_case.Get_All_Measurement
}

func NewGetAllHeartRateController(useCase *use_case.Get_All_Measurement) *GetAllHeartRateController {
	return &GetAllHeartRateController{useCase: useCase}
}

func (gt *GetAllHeartRateController) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID inválido"})
		return
	}

	data, err := gt.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los registros", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
