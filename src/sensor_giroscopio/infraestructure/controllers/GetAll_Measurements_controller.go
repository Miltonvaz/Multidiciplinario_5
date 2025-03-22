package controllers

import (
	"Multidiciplinario/src/sensor_giroscopio/application/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllGyroscopeSensorController struct {
	useCase *use_case.GetAllGyroscopeSensors
}

func NewGetAllGyroscopeSensorController(useCase *use_case.GetAllGyroscopeSensors) *GetAllGyroscopeSensorController {
	return &GetAllGyroscopeSensorController{useCase: useCase}
}

func (gt *GetAllGyroscopeSensorController) Execute(ctx *gin.Context) {
	data, err := gt.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los registros de giroscopio", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
