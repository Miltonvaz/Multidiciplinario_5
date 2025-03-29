package routes_h

import (
	"Multidiciplinario/src/core/security"
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
		api.GET("/heart-rate/:id/:user_id", security.JWTMiddleware(), getHeartRateByIDController.Execute)
		api.GET("/all/:user_id", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:user_id", security.JWTMiddleware(), deleteController.Execute)
		api.GET("/heart-rate/average/:user_id", security.JWTMiddleware(), getAverageHeartRateController.Execute)
		api.GET("/heart-rate/latest/:user_id", security.JWTMiddleware(), getLatestMeasurementController.Execute)
	}
}
