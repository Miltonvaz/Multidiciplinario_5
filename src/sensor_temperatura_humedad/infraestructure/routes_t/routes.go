package routes_t

import (
	"Multidiciplinario/src/sensor_temperatura_humedad/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	createController *controllers.Create_TemperatureAndHumidity_C,
	getTemperatureByIDController *controllers.GetTemperatureByIDController,
	getAllController *controllers.GetAllTemperatureAndHumidityController,
	deleteController *controllers.DeleteTemperatureAndHumidityController,
	getAverageTemperatureController *controllers.GetAverageTemperatureController,
	getAverageHumidityController *controllers.GetAverageHumidityController,
	getLatestMeasurementController *controllers.GetLatestMeasurementController,
) {
	api := router.Group("/api/sensor-TH")
	{
		api.POST("/create", createController.Execute)
		api.GET("/temperature/:id", getTemperatureByIDController.Execute)
		api.GET("/all", getAllController.Execute)
		api.DELETE("/delete/:id", deleteController.Execute)
		api.GET("/temperature/average", getAverageTemperatureController.Execute)
		api.GET("/humidity/average", getAverageHumidityController.Execute)
		api.GET("/temperature/latest", getLatestMeasurementController.Execute)
	}
}
