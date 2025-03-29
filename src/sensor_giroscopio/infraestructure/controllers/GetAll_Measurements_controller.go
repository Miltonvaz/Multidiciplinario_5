package controllers

import (
	"Multidiciplinario/src/sensor_giroscopio/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAllGyroscopeSensorController struct {
	useCase *use_case.GetAllGyroscopeSensors
}

func NewGetAllGyroscopeSensorController(useCase *use_case.GetAllGyroscopeSensors) *GetAllGyroscopeSensorController {
	return &GetAllGyroscopeSensorController{useCase: useCase}
}

func (gt *GetAllGyroscopeSensorController) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID inválido"})
		return
	}

	data, err := gt.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los registros de giroscopio", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
