package routes_t

import (
	"Multidiciplinario/src/core/security"
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
		api.GET("/temperature/:id/:userID", security.JWTMiddleware(), getTemperatureByIDController.Execute)
		api.GET("/all/:userID", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:userID", security.JWTMiddleware(), deleteController.Execute)
		api.GET("/temperature/average/:userID", security.JWTMiddleware(), getAverageTemperatureController.Execute)
		api.GET("/temperature/latest/:userID", security.JWTMiddleware(), getLatestMeasurementController.Execute)
	}
}
