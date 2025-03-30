package dependencies_g

import (
	"Multidiciplinario/src/core"
	"Multidiciplinario/src/sensor_giroscopio/application/repositories"
	"Multidiciplinario/src/sensor_giroscopio/application/use_case"
	"Multidiciplinario/src/sensor_giroscopio/domain"
	"Multidiciplinario/src/sensor_giroscopio/infraestructure/adapters"
	adapters_g "Multidiciplinario/src/sensor_giroscopio/infraestructure/adapters/adapaters_g"
	"Multidiciplinario/src/sensor_giroscopio/infraestructure/controllers"
	"log"
)

func Init(pool *core.Conn_MySQL) (
	*controllers.CreateGyroscopeSensorController,
	*controllers.GetGyroscopeSensorByIDController,
	*controllers.GetAllGyroscopeSensorController,
	*controllers.DeleteGyroscopeSensorController,
	*controllers.GetAverageGyroscopeSensorController,
	*controllers.GetLatestGyroscopeSensorController,
	domain.IGyroscopeSensor,
	*repositories.ServiceNotification,
	*adapters_g.MQTTAdapter,
	error,
) {
	// Initialize repository using the DB connection pool
	repository := adapters.NewMySQL(pool.DB)

	// Initialize RabbitMQ adapter
	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	// Initialize service notification (using the RabbitMQ adapter)
	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	// Initialize MQTT adapter
	mqttAdapter, err := adapters_g.NewMQTTAdapter(repository, serviceNotification)
	if err != nil {
		log.Printf("Error initializing MQTT: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	// Initialize use cases
	createGyroscopeUseCase := use_case.NewCreate_GyroscopeSensor(repository, serviceNotification)
	getGyroscopeByIDUseCase := use_case.NewGetGyroscopeSensorByID(repository)
	getAllGyroscopeSensorsUseCase := use_case.NewGetAllGyroscopeSensors(repository)
	getAverageGyroscopeUseCase := use_case.NewGetAverageGyroscopeData(repository)
	getLatestGyroscopeUseCase := use_case.NewGet_Latest_Gyroscope(repository)

	// Initialize controllers using use cases
	createGyroscopeController := controllers.NewCreateGyroscopeSensorController(createGyroscopeUseCase)
	getGyroscopeByIDController := controllers.NewGetGyroscopeSensorByIDController(getGyroscopeByIDUseCase)
	getAllGyroscopeController := controllers.NewGetAllGyroscopeSensorController(getAllGyroscopeSensorsUseCase)
	getAverageGyroscopeController := controllers.NewGetAverageGyroscopeSensorController(getAverageGyroscopeUseCase)
	getLatestGyroscopeController := controllers.NewGetLatestGyroscopeSensorController(getLatestGyroscopeUseCase)

	// Return all the controllers and dependencies
	return createGyroscopeController, getGyroscopeByIDController, getAllGyroscopeController, nil, getAverageGyroscopeController, getLatestGyroscopeController, repository, serviceNotification, mqttAdapter, nil
}
