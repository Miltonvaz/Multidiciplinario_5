package controllers

import (
	"Multidiciplinario/src/sensor_giroscopio/application/use_case"
	entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"

	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateGyroscopeSensorController struct {
	UseCase *use_case.Create_GyroscopeSensor
}

func NewCreateGyroscopeSensorController(useCase *use_case.Create_GyroscopeSensor) *CreateGyroscopeSensorController {
	return &CreateGyroscopeSensorController{UseCase: useCase}
}

func (c *CreateGyroscopeSensorController) Execute(ctx *gin.Context) {
	var sensor entities.GyroscopeSensor

	if err := ctx.ShouldBindJSON(&sensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	createdSensor, err := c.UseCase.Execute(sensor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar los datos"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente", "data": createdSensor})
}
