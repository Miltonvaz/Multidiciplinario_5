package controllers

import (
	"Multidiciplinario/src/sensor_luz/application/use_case"
	"Multidiciplinario/src/sensor_luz/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_LightLDR_C struct {
	UseCase *use_case.Create_LightLDR
}

func NewCreate_LightLDR_C(useCase *use_case.Create_LightLDR) *Create_LightLDR_C {
	return &Create_LightLDR_C{UseCase: useCase}
}

func (c *Create_LightLDR_C) Execute(ctx *gin.Context) {
	var sensor entities.LightSensorLDR

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
