package dependencies_t

import (
	"Multidiciplinario/src/core"
	"Multidiciplinario/src/sensor_temperatura_corporal/application/repositories"
	"Multidiciplinario/src/sensor_temperatura_corporal/application/use_case"
	"Multidiciplinario/src/sensor_temperatura_corporal/infraestructure/adapters"
	"Multidiciplinario/src/sensor_temperatura_corporal/infraestructure/adapters/adapter_m"
	"Multidiciplinario/src/sensor_temperatura_corporal/infraestructure/controllers"
	"log"
)

func Init(pool *core.Conn_MySQL) (
	*controllers.Create_BodyTemperature_C,
	*controllers.GetTemperatureByIDController,
	*controllers.GetAllTemperatureAndHumidityController,
	*controllers.DeleteBodyTemperatureController,
	*controllers.GetAverageTemperatureController,
	*controllers.GetLatestMeasurementController,
	*repositories.ServiceNotification,
	*adapter_m.MQTTAdapter,
	error,
) {

	repository := adapters.NewMySQL(pool.DB)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	mqttAdapter, err := adapter_m.NewMQTTAdapter(repository, serviceNotification)
	if err != nil {
		log.Printf("Error initializing MQTT adapter for body temperature: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	createUseCase := use_case.NewCreate_BodyTemperature(repository, serviceNotification)
	getTemperatureByIDUseCase := use_case.NewGetMeasurementByID(repository)
	getAllUseCase := use_case.NewGet_All(repository)
	deleteUseCase := use_case.NewDelete_TemperatureAndHumidity(repository)
	getAverageTemperatureUseCase := use_case.NewGetAverageTemperature(repository)
	getLatestMeasurementUseCase := use_case.NewGet_Latest(repository)

	createController := controllers.NewCreate_TemperatureAndHumidity_C(createUseCase)
	getTemperatureByIDController := controllers.NewGetTemperatureByIDController(getTemperatureByIDUseCase)
	getAllController := controllers.NewGetAllTemperatureAndHumidityController(getAllUseCase)
	deleteController := controllers.NewDeleteTemperatureAndHumidityController(deleteUseCase)
	getAverageTemperatureController := controllers.NewGetAverageTemperatureController(getAverageTemperatureUseCase)
	getLatestMeasurementController := controllers.NewGetLatestMeasurementController(getLatestMeasurementUseCase)

	return createController,
		getTemperatureByIDController,
		getAllController,
		deleteController,
		getAverageTemperatureController,
		getLatestMeasurementController,
		serviceNotification,
		mqttAdapter,
		nil
}
