package dependencies_h

import (
	"Multidiciplinario/src/core"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/repositories"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/use_case"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/infraestructure/adapters"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/infraestructure/adapters/adapter_h"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/infraestructure/controllers"
	"log"
)

func Init() (
	*controllers.Create_HeartRate_C,
	*controllers.GetHeartRateByIDController,
	*controllers.GetAllHeartRateController,
	*controllers.DeleteHeartRateController,
	*controllers.GetAverageHeartRateController,
	*controllers.GetLatestMeasurementController,
	*repositories.ServiceNotification,
	*adapter_h.MQTTAdapter,
	error,
) {

	pool := core.GetDBPool()
	repository := adapters.NewMySQL(pool.DB)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	mqttAdapter, err := adapter_h.NewMQTTAdapter(repository, serviceNotification)
	if err != nil {
		log.Printf("Error initializing MQTT adapter for heart rate: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	createUseCase := use_case.NewCreate_HeartRate(repository, serviceNotification)
	getHeartRateByIDUseCase := use_case.NewGetMeasurementByID(repository)
	getAllUseCase := use_case.NewGet_All(repository)
	deleteUseCase := use_case.NewDelete_HeartRate(repository)
	getAverageHeartRateUseCase := use_case.NewGetAverageHeartRate(repository)
	getLatestMeasurementUseCase := use_case.NewGet_Latest(repository)

	createController := controllers.NewCreate_HeartRate_C(createUseCase)
	getHeartRateByIDController := controllers.NewGetHeartRateByIDController(getHeartRateByIDUseCase)
	getAllController := controllers.NewGetAllHeartRateController(getAllUseCase)
	deleteController := controllers.NewDeleteHeartRateController(deleteUseCase)
	getAverageHeartRateController := controllers.NewGetAverageHeartRateController(getAverageHeartRateUseCase)
	getLatestMeasurementController := controllers.NewGetLatestMeasurementController(getLatestMeasurementUseCase)

	return createController,
		getHeartRateByIDController,
		getAllController,
		deleteController,
		getAverageHeartRateController,
		getLatestMeasurementController,
		serviceNotification,
		mqttAdapter,
		nil
}
