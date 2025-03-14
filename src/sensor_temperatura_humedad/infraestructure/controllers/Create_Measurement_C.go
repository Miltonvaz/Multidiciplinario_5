package controllers

import (
	"Multidiciplinario/src/sensor_temperatura_humedad/application/use_case"
	"Multidiciplinario/src/sensor_temperatura_humedad/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_TemperatureAndHumidity_C struct {
	UseCase *use_case.Create_TemperatureAndHumidity
}

func NewCreate_TemperatureAndHumidity_C(useCase *use_case.Create_TemperatureAndHumidity) *Create_TemperatureAndHumidity_C {
	return &Create_TemperatureAndHumidity_C{UseCase: useCase}
}

func (c *Create_TemperatureAndHumidity_C) Execute(ctx *gin.Context) {
	var sensor entities.TemperatureAndHumidity

	// Validate the incoming data
	if err := ctx.ShouldBindJSON(&sensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Execute the use case and handle the return values
	createdSensor, err := c.UseCase.Execute(sensor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar los datos"})
		return
	}

	// Respond with the created sensor data
	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente", "data": createdSensor})
}
