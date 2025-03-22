package controllers

import (
	"Multidiciplinario/src/sensor_giroscopio/application/use_case"
	entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetGyroscopeSensorByIDController struct {
	useCase *use_case.GetGyroscopeSensorByID
}

func NewGetGyroscopeSensorByIDController(useCase *use_case.GetGyroscopeSensorByID) *GetGyroscopeSensorByIDController {
	return &GetGyroscopeSensorByIDController{useCase: useCase}
}

func (gh *GetGyroscopeSensorByIDController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	measurement, err := gh.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos de giroscopio", "details": err.Error()})
		return
	}
	if (measurement == entities.GyroscopeSensor{}) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontró el registro con el ID proporcionado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": measurement,
	})
}
