package routes_t

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	createController *controllers.Create_BodyTemperature_C,
	getTemperatureByIDController *controllers.GetTemperatureByIDController,
	getAllController *controllers.GetAllTemperatureAndHumidityController,
	deleteController *controllers.DeleteBodyTemperatureController,
	getAverageTemperatureController *controllers.GetAverageTemperatureController,
	getLatestMeasurementController *controllers.GetLatestMeasurementController,
) {
	api := router.Group("/api/sensor-TH")
	{
		api.POST("/create", createController.Execute)
		api.GET("/temperature/:id", getTemperatureByIDController.Execute)
		api.GET("/all", getAllController.Execute)
		api.DELETE("/delete/:id", deleteController.Execute)
		api.GET("/temperature/average", getAverageTemperatureController.Execute)
		api.GET("/temperature/latest", getLatestMeasurementController.Execute)
	}
}
