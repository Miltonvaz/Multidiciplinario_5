package controllers

import (
	"Multidiciplinario/src/sensor_luz/application/use_case"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAllLightLDRController struct {
	useCase *use_case.Get_All_Measurement
}

func NewGetAllLightLDRController(useCase *use_case.Get_All_Measurement) *GetAllLightLDRController {
	return &GetAllLightLDRController{useCase: useCase}
}
func (gt *GetAllLightLDRController) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)

	data, err := gt.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los registros", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
