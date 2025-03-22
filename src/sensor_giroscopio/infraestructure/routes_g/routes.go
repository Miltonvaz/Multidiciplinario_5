package routes_g

import (
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
		api.GET("/gyroscope/:id", getGyroscopeSensorByIDController.Execute)
		api.GET("/all", getAllController.Execute)
		api.DELETE("/delete/:id", deleteController.Execute)
		api.GET("/gyroscope/average", getAverageGyroscopeSensorController.Execute)
		api.GET("/gyroscope/latest", getLatestMeasurementController.Execute)
	}
}
