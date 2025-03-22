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
	"Multidiciplinario/src/users/infraestructure/dependencies_u"
	"Multidiciplinario/src/users/infraestructure/routes_u"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func startServer(router *gin.Engine) error {
	serverErr := router.Run(":8086")
	if serverErr != nil {
		log.Printf("Error en el servidor: %v", serverErr)
	}
	return serverErr
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// Inicialización de las dependencias y controladores
	createTemperatureController, getTemperatureByIDController, getAllTemperatureController, deleteTemperatureController, getAverageTemperatureController, getLatestTemperatureMeasurementController, _, _, temperatureErr := dependencies_t.Init()
	if temperatureErr != nil {
		log.Fatalf("Error al inicializar las dependencias de temperatura: %v", temperatureErr)
		return
	}

	createHeartRateController, getHeartRateByIDController, getAllHeartRateController, deleteHeartRateController, getAverageHeartRateController, getLatestHeartRateMeasurementController, _, _, heartRateErr := dependencies_h.Init()
	if heartRateErr != nil {
		log.Fatalf("Error al inicializar las dependencias de ritmo cardíaco: %v", heartRateErr)
		return
	}

	createLightController, getLightByIDController, getAllLightController, deleteLightController, getAverageLightController, getLatestLightMeasurementController, _, _, _, lightErr := dependencies_l.Init()
	if lightErr != nil {
		log.Fatalf("Error al inicializar las dependencias de luz: %v", lightErr)
		return
	}

	createGyroscopeController, getGyroscopeByIDController, getAllGyroscopeController, deleteGyroscopeController, getAverageGyroscopeController, getLatestGyroscopeMeasurementController, _, _, _, gyroscopeErr := dependencies_g.Init()
	if gyroscopeErr != nil {
		log.Fatalf("Error al inicializar las dependencias del giroscopio: %v", gyroscopeErr)
		return
	}

	createUserController, viewUserController, editUserController, deleteUserController, viewByIdUserController, loginController, userErr := dependencies_u.Init()
	if userErr != nil {
		log.Fatalf("Error al inicializar las dependencias de usuario: %v", userErr)
		return
	}

	// Rutas para cada uno de los servicios
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

	routes_u.RegisterClientRoutes(
		router,
		createUserController,
		viewUserController,
		editUserController,
		deleteUserController,
		viewByIdUserController,
		loginController,
	)

	// Inicia el servidor
	if err := startServer(router); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
