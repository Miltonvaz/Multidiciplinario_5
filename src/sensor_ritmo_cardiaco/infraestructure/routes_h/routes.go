package routes_h

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterHeartRateRoutes(
	router *gin.Engine,
	createController *controllers.Create_HeartRate_C,
	getHeartRateByIDController *controllers.GetHeartRateByIDController,
	getAllController *controllers.GetAllHeartRateController,
	deleteController *controllers.DeleteHeartRateController,
	getAverageHeartRateController *controllers.GetAverageHeartRateController,
	getLatestMeasurementController *controllers.GetLatestMeasurementController,
) {
	api := router.Group("/api/sensor-heart-rate")
	{
		api.POST("/create", createController.Execute)
		api.GET("/heart-rate/:id", getHeartRateByIDController.Execute)
		api.GET("/all", getAllController.Execute)
		api.DELETE("/delete/:id", deleteController.Execute)
		api.GET("/heart-rate/average", getAverageHeartRateController.Execute)
		api.GET("/heart-rate/latest", getLatestMeasurementController.Execute)
	}
}
