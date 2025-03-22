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

	createUseCase := use_case.NewCreate_GyroscopeSensor(repository, serviceNotification)
	getGyroscopeSensorByIDUseCase := use_case.NewGetGyroscopeSensorByID(repository)
	getAllUseCase := use_case.NewGetAllGyroscopeSensors(repository)
	deleteUseCase := use_case.NewDelete_GyroscopeSensor(repository)
	getAverageGyroscopeSensorUseCase := use_case.NewGetAverageGyroscopeData(repository)
	getLatestMeasurementUseCase := use_case.NewGet_Latest_Gyroscope(repository)

	createController := controllers.NewCreateGyroscopeSensorController(createUseCase)
	getGyroscopeSensorByIDController := controllers.NewGetGyroscopeSensorByIDController(getGyroscopeSensorByIDUseCase)
	getAllController := controllers.NewGetAllGyroscopeSensorController(getAllUseCase)
	deleteController := controllers.NewDeleteGyroscopeSensorController(deleteUseCase)
	getAverageGyroscopeSensorController := controllers.NewGetAverageGyroscopeSensorController(getAverageGyroscopeSensorUseCase)
	getLatestMeasurementController := controllers.NewGetLatestGyroscopeSensorController(getLatestMeasurementUseCase)

	return createController,
		getGyroscopeSensorByIDController,
		getAllController,
		deleteController,
		getAverageGyroscopeSensorController,
		getLatestMeasurementController,
		repository,
		serviceNotification,
		mqttAdapter,
		nil
}
