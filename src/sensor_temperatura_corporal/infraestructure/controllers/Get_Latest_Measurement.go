package controllers

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/application/use_case"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetLatestMeasurementController struct {
	useCase *use_case.Get_Latest
}

func NewGetLatestMeasurementController(useCase *use_case.Get_Latest) *GetLatestMeasurementController {
	return &GetLatestMeasurementController{useCase: useCase}
}

func (gl *GetLatestMeasurementController) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID inválido"})
		return
	}

	data, err := gl.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la última medición", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
