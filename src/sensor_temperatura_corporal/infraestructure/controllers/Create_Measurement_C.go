package controllers

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/application/use_case"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_BodyTemperature_C struct {
	UseCase *use_case.Create_BodyTemperature
}

func NewCreate_TemperatureAndHumidity_C(useCase *use_case.Create_BodyTemperature) *Create_BodyTemperature_C {
	return &Create_BodyTemperature_C{UseCase: useCase}
}

func (c *Create_BodyTemperature_C) Execute(ctx *gin.Context) {
	var sensor entities.BodyTemperature

	// Deserialize the received JSON data
	if err := ctx.ShouldBindJSON(&sensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Execute use case to save the data
	createdSensor, err := c.UseCase.Execute(sensor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Data saved successfully", "data": createdSensor})
}
