package controllers

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteHeartRateController struct {
	useCase *use_case.Delete_HeartRate
}

func NewDeleteHeartRateController(useCase *use_case.Delete_HeartRate) *DeleteHeartRateController {
	return &DeleteHeartRateController{useCase: useCase}
}

func (ct *DeleteHeartRateController) Execute(ctx *gin.Context) {
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
