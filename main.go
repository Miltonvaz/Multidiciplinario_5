package main

import (
	"Multidiciplinario/src/core"
	"Multidiciplinario/src/sensor_giroscopio/infraestructure/dependencies_g"
	"Multidiciplinario/src/sensor_giroscopio/infraestructure/routes_g"
	"Multidiciplinario/src/sensor_luz/infraestructure/dependencies_l"
	"Multidiciplinario/src/sensor_luz/infraestructure/routes_l"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/infraestructure/dependencies_h"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/infraestructure/routes_h"
	"Multidiciplinario/src/sensor_temperatura_corporal/infraestructure/dependencies_t"
	"Multidiciplinario/src/sensor_temperatura_corporal/infraestructure/routes_t"
	"Multidiciplinario/src/users/infraestructure/dependencies_u"
	"Multidiciplinario/src/users/infraestructure/routes_u"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func startServer() {
	for {
		log.Println("Iniciando servidor...")
		router := gin.Default()

		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
		config.AllowHeaders = []string{"Authorization", "Content-Type"}
		config.ExposeHeaders = []string{"Content-Length", "Authorization"}
		config.MaxAge = 12 * time.Hour

		router.Use(cors.New(config))

		if err := initializeDependencies(router); err != nil {
			log.Fatalf("Error al inicializar dependencias: %v", err)
			return
		}
		go func() {
			if err := router.Run(":8080"); err != nil {
				log.Printf("Error en el servidor: %v", err)
			}
		}()
		time.Sleep(3 * time.Minute)
		log.Println("Reiniciando servidor...")
	}
}

func initializeDependencies(router *gin.Engine) error {

	pool := core.GetDBPool()

	createTemperatureController, getTemperatureByIDController, getAllTemperatureController, deleteTemperatureController, getAverageTemperatureController, getLatestTemperatureMeasurementController, _, _, temperatureErr := dependencies_t.Init(pool)
	if temperatureErr != nil {
		return temperatureErr
	}

	createHeartRateController, getHeartRateByIDController, getAllHeartRateController, deleteHeartRateController, getAverageHeartRateController, getLatestHeartRateMeasurementController, _, _, heartRateErr := dependencies_h.Init(pool)
	if heartRateErr != nil {
		return heartRateErr
	}

	createLightController, getLightByIDController, getAllLightController, deleteLightController, getAverageLightController, getLatestLightMeasurementController, _, _, _, lightErr := dependencies_l.Init(pool)
	if lightErr != nil {
		return lightErr
	}

	createGyroscopeController, getGyroscopeByIDController, getAllGyroscopeController, deleteGyroscopeController, getAverageGyroscopeController, getLatestGyroscopeMeasurementController, _, _, _, gyroscopeErr := dependencies_g.Init(pool)
	if gyroscopeErr != nil {
		return gyroscopeErr
	}

	createUserController, viewUserController, editUserController, deleteUserController, viewByIdUserController, loginController, userErr := dependencies_u.Init(pool)
	if userErr != nil {
		return userErr
	}

	// Register routes
	routes_t.RegisterRoutes(router, createTemperatureController, getTemperatureByIDController, getAllTemperatureController, deleteTemperatureController, getAverageTemperatureController, getLatestTemperatureMeasurementController)
	routes_h.RegisterHeartRateRoutes(router, createHeartRateController, getHeartRateByIDController, getAllHeartRateController, deleteHeartRateController, getAverageHeartRateController, getLatestHeartRateMeasurementController)
	routes_l.RegisterLightSensorRoutes(router, createLightController, getLightByIDController, getAllLightController, deleteLightController, getAverageLightController, getLatestLightMeasurementController)
	routes_g.RegisterGyroscopeSensorRoutes(router, createGyroscopeController, getGyroscopeByIDController, getAllGyroscopeController, deleteGyroscopeController, getAverageGyroscopeController, getLatestGyroscopeMeasurementController)
	routes_u.RegisterClientRoutes(router, createUserController, viewUserController, editUserController, deleteUserController, viewByIdUserController, loginController)

	return nil
}

func main() {
	startServer()
}
