package controllers

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/use_case"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetHeartRateByIDController struct {
	useCase *use_case.GetMeasurementByID
}

func NewGetHeartRateByIDController(useCase *use_case.GetMeasurementByID) *GetHeartRateByIDController {
	return &GetHeartRateByIDController{useCase: useCase}
}

func (gh *GetHeartRateByIDController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userIDStr := ctx.Param("user_id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID inválido"})
		return
	}

	measurement, err := gh.useCase.Execute(id, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la medición", "details": err.Error()})
		return
	}
	if (measurement == entities.HeartRate{}) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontró el registro con el ID proporcionado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": measurement})
}
