package controllers

import (
	"Multidiciplinario/src/sensor_luz/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetLatestMeasurementController struct {
	useCase *use_case.Get_Latest
}

func NewGetLatestMeasurementController(useCase *use_case.Get_Latest) *GetLatestMeasurementController {
	return &GetLatestMeasurementController{useCase: useCase}
}

func (gl *GetLatestMeasurementController) Execute(ctx *gin.Context) {
	data, err := gl.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la última medición", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
