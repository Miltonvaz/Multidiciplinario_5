package controllers

import (
	"Multidiciplinario/src/sensor_temperatura_humedad/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteTemperatureAndHumidityController struct {
	useCase *use_case.Delete_TemperatureAndHumidity
}

func NewDeleteTemperatureAndHumidityController(useCase *use_case.Delete_TemperatureAndHumidity) *DeleteTemperatureAndHumidityController {
	return &DeleteTemperatureAndHumidityController{useCase: useCase}
}

// Ejecuta la lógica para eliminar un sensor de temperatura y humedad
func (ct *DeleteTemperatureAndHumidityController) Execute(ctx *gin.Context) {
	// Obtener el ID del sensor desde la URL como string
	idStr := ctx.Param("id")

	// Convertir el ID de string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Si hay un error de conversión, devolver error 400
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar al caso de uso para eliminar el sensor
	if err := ct.useCase.Execute(id); err != nil {
		// Si ocurre un error, devolver mensaje de error
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar los datos"})
		return
	}

	// Si la eliminación es exitosa, devolver respuesta exitosa
	ctx.JSON(http.StatusOK, gin.H{"message": "Datos eliminados correctamente"})
}
