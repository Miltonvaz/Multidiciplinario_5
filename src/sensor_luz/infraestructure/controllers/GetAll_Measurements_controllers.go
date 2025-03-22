package controllers

import (
	"Multidiciplinario/src/sensor_luz/application/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllLightLDRController struct {
	useCase *use_case.Get_All_Measurement
}

func NewGetAllLightLDRController(useCase *use_case.Get_All_Measurement) *GetAllLightLDRController {
	return &GetAllLightLDRController{useCase: useCase}
}
func (gt *GetAllLightLDRController) Execute(ctx *gin.Context) {
	data, err := gt.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los registros", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
