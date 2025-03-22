package routes_l

import (
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
		api.GET("/light/:id", getLightLDRByIDController.Execute)
		api.GET("/all", getAllController.Execute)
		api.DELETE("/delete/:id", deleteController.Execute)
		api.GET("/light/average", getAverageLightLDRController.Execute)
		api.GET("/light/latest", getLatestMeasurementController.Execute)
	}
}
