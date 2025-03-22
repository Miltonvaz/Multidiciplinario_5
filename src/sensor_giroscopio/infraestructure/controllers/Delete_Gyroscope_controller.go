package controllers

import (
	"Multidiciplinario/src/sensor_giroscopio/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteGyroscopeSensorController struct {
	useCase *use_case.Delete_GyroscopeSensor
}

func NewDeleteGyroscopeSensorController(useCase *use_case.Delete_GyroscopeSensor) *DeleteGyroscopeSensorController {
	return &DeleteGyroscopeSensorController{useCase: useCase}
}

func (ct *DeleteGyroscopeSensorController) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ct.useCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar los datos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Datos eliminados correctamente"})
}
