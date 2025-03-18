package main

import (
	"Multidiciplinario/src/sensor_temperatura_humedad/infraestructure/adapters/adapter_m"
	"Multidiciplinario/src/sensor_temperatura_humedad/infraestructure/dependencies_t"
	"Multidiciplinario/src/sensor_temperatura_humedad/infraestructure/routes_t"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	createController, getTemperatureByIDController, getAllController, deleteController, getAverageTemperatureController, getAverageHumidityController, getLatestMeasurementController, repository, serviceNotification, err := dependencies_t.Init()
	if err != nil {
		log.Fatalf("Error initializing dependencies: %v", err)
		return
	}

	routes_t.RegisterRoutes(router, createController, getTemperatureByIDController, getAllController, deleteController, getAverageTemperatureController, getAverageHumidityController, getLatestMeasurementController)

	_, err = adapter_m.NewMQTTAdapter(repository, serviceNotification)
	if err != nil {
		log.Fatalf("Error initializing MQTT adapter: %v", err)
		return
	}

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
