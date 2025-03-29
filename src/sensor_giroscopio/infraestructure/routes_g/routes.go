package routes_g

import (
	"Multidiciplinario/src/core/security"
	"Multidiciplinario/src/sensor_giroscopio/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterGyroscopeSensorRoutes(
	router *gin.Engine,
	createController *controllers.CreateGyroscopeSensorController,
	getGyroscopeSensorByIDController *controllers.GetGyroscopeSensorByIDController,
	getAllController *controllers.GetAllGyroscopeSensorController,
	deleteController *controllers.DeleteGyroscopeSensorController,
	getAverageGyroscopeSensorController *controllers.GetAverageGyroscopeSensorController,
	getLatestMeasurementController *controllers.GetLatestGyroscopeSensorController,
) {
	api := router.Group("/api/sensor-gyroscope")
	{
		api.POST("/create", createController.Execute)
		api.GET("/gyroscope/:id/:user_id", security.JWTMiddleware(), getGyroscopeSensorByIDController.Execute)
		api.GET("/all/:user_id", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:user_id", security.JWTMiddleware(), deleteController.Execute)
		api.GET("/gyroscope/average/:user_id", security.JWTMiddleware(), getAverageGyroscopeSensorController.Execute)
		api.GET("/gyroscope/latest/:user_id", security.JWTMiddleware(), getLatestMeasurementController.Execute)
	}
}
