package controllers

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllHeartRateController struct {
	useCase *use_case.Get_All_Measurement
}

func NewGetAllHeartRateController(useCase *use_case.Get_All_Measurement) *GetAllHeartRateController {
	return &GetAllHeartRateController{useCase: useCase}
}
func (gt *GetAllHeartRateController) Execute(ctx *gin.Context) {
	data, err := gt.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los registros", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
