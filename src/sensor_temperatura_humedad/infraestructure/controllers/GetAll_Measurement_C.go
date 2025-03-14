package controllers

import (
	"Multidiciplinario/src/sensor_temperatura_humedad/application/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllTemperatureAndHumidityController struct {
	useCase *use_case.Get_All_Measurement
}

func NewGetAllTemperatureAndHumidityController(useCase *use_case.Get_All_Measurement) *GetAllTemperatureAndHumidityController {
	return &GetAllTemperatureAndHumidityController{useCase: useCase}
}
func (gt *GetAllTemperatureAndHumidityController) Execute(ctx *gin.Context) {
	data, err := gt.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los registros", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
