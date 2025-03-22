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

func Init() (
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
	pool := core.GetDBPool()
	repository := adapters.NewMySQL(pool.DB)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	mqttAdapter, err := adapters_g.NewMQTTAdapter(repository, serviceNotification)
	if err != nil {
		log.Printf("Error initializing MQTT: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	createGyroscopeUseCase := use_case.NewCreate_GyroscopeSensor(repository, serviceNotification)
	getGyroscopeByIDUseCase := use_case.NewGetGyroscopeSensorByID(repository)
	getAllGyroscopeSensorsUseCase := use_case.NewGetAllGyroscopeSensors(repository)
	getAverageGyroscopeUseCase := use_case.NewGetAverageGyroscopeData(repository)
	getLatestGyroscopeUseCase := use_case.NewGet_Latest_Gyroscope(repository)

	createGyroscopeController := controllers.NewCreateGyroscopeSensorController(createGyroscopeUseCase)
	getGyroscopeByIDController := controllers.NewGetGyroscopeSensorByIDController(getGyroscopeByIDUseCase)
	getAllGyroscopeController := controllers.NewGetAllGyroscopeSensorController(getAllGyroscopeSensorsUseCase)
	getAverageGyroscopeController := controllers.NewGetAverageGyroscopeSensorController(getAverageGyroscopeUseCase)
	getLatestGyroscopeController := controllers.NewGetLatestGyroscopeSensorController(getLatestGyroscopeUseCase)

	return createGyroscopeController, getGyroscopeByIDController, getAllGyroscopeController, nil, getAverageGyroscopeController, getLatestGyroscopeController, repository, serviceNotification, mqttAdapter, nil
}
