package routes_l

import (
	"Multidiciplinario/src/core/security"
	"Multidiciplinario/src/sensor_luz/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterLightSensorRoutes(
	router *gin.Engine,
	createController *controllers.Create_LightLDR_C,
	getLightLDRByIDController *controllers.GetLightLDRByIDController,
	getAllController *controllers.GetAllLightLDRController,
	deleteController *controllers.DeleteLightLDRController,
	getAverageLightLDRController *controllers.GetAverageLightLDRController,
	getLatestMeasurementController *controllers.GetLatestMeasurementController,
) {
	api := router.Group("/api/sensor-light")
	{
		api.POST("/create", createController.Execute)
		api.GET("/light/:id/:user_id", security.JWTMiddleware(), getLightLDRByIDController.Execute)
		api.GET("/all/:user_id", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:user_id", security.JWTMiddleware(), deleteController.Execute)
		api.GET("/light/average/:user_id", security.JWTMiddleware(), getAverageLightLDRController.Execute)
		api.GET("/light/latest/:user_id", security.JWTMiddleware(), getLatestMeasurementController.Execute)
	}
}
