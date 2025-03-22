package controllers

import (
	"Multidiciplinario/src/sensor_giroscopio/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetLatestGyroscopeSensorController struct {
	useCase *use_case.Get_Latest_Gyroscope
}

func NewGetLatestGyroscopeSensorController(useCase *use_case.Get_Latest_Gyroscope) *GetLatestGyroscopeSensorController {
	return &GetLatestGyroscopeSensorController{useCase: useCase}
}

func (gl *GetLatestGyroscopeSensorController) Execute(ctx *gin.Context) {
	data, err := gl.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la última medición de giroscopio", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
