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
		api.GET("/light/:id/:user_id", getLightLDRByIDController.Execute)
		api.GET("/all/:user_id", getAllController.Execute)
		api.DELETE("/delete/:id/:user_id", deleteController.Execute)
		api.GET("/light/average/:user_id", getAverageLightLDRController.Execute)
		api.GET("/light/latest/:user_id", getLatestMeasurementController.Execute)
	}
}
