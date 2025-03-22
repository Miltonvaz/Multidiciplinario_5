package main

import (
	"Multidiciplinario/src/sensor_giroscopio/infraestructure/dependencies_g"
	"Multidiciplinario/src/sensor_giroscopio/infraestructure/routes_g"
	"Multidiciplinario/src/sensor_luz/infraestructure/dependencies_l"
	"Multidiciplinario/src/sensor_luz/infraestructure/routes_l"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/infraestructure/dependencies_h"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/infraestructure/routes_h"
	"Multidiciplinario/src/sensor_temperatura_corporal/infraestructure/dependencies_t"
	"Multidiciplinario/src/sensor_temperatura_corporal/infraestructure/routes_t"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	createTemperatureController, getTemperatureByIDController, getAllTemperatureController, deleteTemperatureController, getAverageTemperatureController, getLatestTemperatureMeasurementController, _, _, temperatureErr := dependencies_t.Init()
	if temperatureErr != nil {
		log.Fatalf("Error initializing temperature dependencies: %v", temperatureErr)
		return
	}

	createHeartRateController, getHeartRateByIDController, getAllHeartRateController, deleteHeartRateController, getAverageHeartRateController, getLatestHeartRateMeasurementController, _, _, heartRateErr := dependencies_h.Init()
	if heartRateErr != nil {
		log.Fatalf("Error initializing heart rate dependencies: %v", heartRateErr)
		return
	}

	createLightController, getLightByIDController, getAllLightController, deleteLightController, getAverageLightController, getLatestLightMeasurementController, _, _, _, lightErr := dependencies_l.Init()
	if lightErr != nil {
		log.Fatalf("Error initializing light sensor dependencies: %v", lightErr)
		return
	}

	createGyroscopeController, getGyroscopeByIDController, getAllGyroscopeController, deleteGyroscopeController, getAverageGyroscopeController, getLatestGyroscopeMeasurementController, _, _, _, gyroscopeErr := dependencies_g.Init()
	if gyroscopeErr != nil {
		log.Fatalf("Error initializing gyroscope sensor dependencies: %v", gyroscopeErr)
		return
	}

	routes_t.RegisterRoutes(
		router,
		createTemperatureController,
		getTemperatureByIDController,
		getAllTemperatureController,
		deleteTemperatureController,
		getAverageTemperatureController,
		getLatestTemperatureMeasurementController,
	)

	routes_h.RegisterHeartRateRoutes(
		router,
		createHeartRateController,
		getHeartRateByIDController,
		getAllHeartRateController,
		deleteHeartRateController,
		getAverageHeartRateController,
		getLatestHeartRateMeasurementController,
	)

	routes_l.RegisterLightSensorRoutes(
		router,
		createLightController,
		getLightByIDController,
		getAllLightController,
		deleteLightController,
		getAverageLightController,
		getLatestLightMeasurementController,
	)

	routes_g.RegisterGyroscopeSensorRoutes(
		router,
		createGyroscopeController,
		getGyroscopeByIDController,
		getAllGyroscopeController,
		deleteGyroscopeController,
		getAverageGyroscopeController,
		getLatestGyroscopeMeasurementController,
	)

	serverErr := router.Run(":8080")
	if serverErr != nil {
		log.Fatalf("Error starting server: %v", serverErr)
	}
}
