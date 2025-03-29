package controllers

import (
	"Multidiciplinario/src/sensor_luz/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteLightLDRController struct {
	useCase *use_case.Delete_LightLDR
}

func NewDeleteLightLDRController(useCase *use_case.Delete_LightLDR) *DeleteLightLDRController {
	return &DeleteLightLDRController{useCase: useCase}
}

func (ct *DeleteLightLDRController) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userIDStr := ctx.Param("user_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID inválido"})
		return
	}

	if err := ct.useCase.Execute(id, userID); err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar los datos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Datos eliminados correctamente"})
}
