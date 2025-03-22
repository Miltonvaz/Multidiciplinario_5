package controllers

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/use_case"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_HeartRate_C struct {
	UseCase *use_case.Create_HeartRate
}

func NewCreate_HeartRate_C(useCase *use_case.Create_HeartRate) *Create_HeartRate_C {
	return &Create_HeartRate_C{UseCase: useCase}
}

func (c *Create_HeartRate_C) Execute(ctx *gin.Context) {
	var sensor entities.HeartRate

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
