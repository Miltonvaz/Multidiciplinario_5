package dependencies_l

import (
	"Multidiciplinario/src/core"
	"Multidiciplinario/src/sensor_luz/application/repositories"
	"Multidiciplinario/src/sensor_luz/application/use_case"
	"Multidiciplinario/src/sensor_luz/domain"
	"Multidiciplinario/src/sensor_luz/infraestructure/adapters"
	"Multidiciplinario/src/sensor_luz/infraestructure/adapters/adapters_l"
	"Multidiciplinario/src/sensor_luz/infraestructure/controllers"
	"log"
)

func Init(pool *core.Conn_MySQL) (
	*controllers.Create_LightLDR_C,
	*controllers.GetLightLDRByIDController,
	*controllers.GetAllLightLDRController,
	*controllers.DeleteLightLDRController,
	*controllers.GetAverageLightLDRController,
	*controllers.GetLatestMeasurementController,
	domain.ILightSensor,
	*repositories.ServiceNotification,
	*adapters_l.MQTTAdapter,
	error,
) {
	repository := adapters.NewMySQL(pool.DB)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	mqttAdapter, err := adapters_l.NewMQTTAdapter(repository, serviceNotification)
	if err != nil {
		log.Printf("Error initializing MQTT: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	createUseCase := use_case.NewCreate_LightLDR(repository, serviceNotification)
	getLightLDRByIDUseCase := use_case.NewGetMeasurementByID(repository)
	getAllUseCase := use_case.NewGet_All(repository)
	deleteUseCase := use_case.NewDelete_LightLDR(repository)
	getAverageLightLDRUseCase := use_case.NewGetAverageLightLDR(repository)
	getLatestMeasurementUseCase := use_case.NewGet_Latest(repository)

	createController := controllers.NewCreate_LightLDR_C(createUseCase)
	getLightLDRByIDController := controllers.NewGetLightLDRByIDController(getLightLDRByIDUseCase)
	getAllController := controllers.NewGetAllLightLDRController(getAllUseCase)
	deleteController := controllers.NewDeleteLightLDRController(deleteUseCase)
	getAverageLightLDRController := controllers.NewGetAverageLightLDRController(getAverageLightLDRUseCase)
	getLatestMeasurementController := controllers.NewGetLatestMeasurementController(getLatestMeasurementUseCase)

	return createController,
		getLightLDRByIDController,
		getAllController,
		deleteController,
		getAverageLightLDRController,
		getLatestMeasurementController,
		repository,
		serviceNotification,
		mqttAdapter,
		nil
}
