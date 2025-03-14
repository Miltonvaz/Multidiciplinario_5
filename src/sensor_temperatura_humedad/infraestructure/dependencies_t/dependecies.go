package dependencies_t

import (
	"Multidiciplinario/src/core"
	"Multidiciplinario/src/sensor_temperatura_humedad/application/repositories"
	"Multidiciplinario/src/sensor_temperatura_humedad/application/use_case"
	"Multidiciplinario/src/sensor_temperatura_humedad/domain"
	"Multidiciplinario/src/sensor_temperatura_humedad/infraestructure/adapters"
	"Multidiciplinario/src/sensor_temperatura_humedad/infraestructure/controllers"
	"log"
)

func Init() (
	*controllers.Create_TemperatureAndHumidity_C,
	*controllers.GetTemperatureByIDController,
	*controllers.GetAllTemperatureAndHumidityController,
	*controllers.DeleteTemperatureAndHumidityController,
	*controllers.GetAverageTemperatureController,
	*controllers.GetAverageHumidityController,
	*controllers.GetLatestMeasurementController,
	domain.ITemperatureAndHumidity,
	*repositories.ServiceNotification, // Añadido aquí
	error,
) {
	// Obtiene la conexión con la base de datos
	pool := core.GetDBPool()
	repository := adapters.NewMySQL(pool.DB)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error inicializando RabbitMQ: %v", err)
	}
	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	createUseCase := use_case.NewCreate_TemperatureAndHumidity(repository, serviceNotification)
	getTemperatureByIDUseCase := use_case.NewGetMeasurementByID(repository)
	getAllUseCase := use_case.NewGet_All(repository)
	deleteUseCase := use_case.NewDelete_TemperatureAndHumidity(repository)
	getAverageTemperatureUseCase := use_case.NewGetAverageTemperature(repository)
	getAverageHumidityUseCase := use_case.NewGetAverageHumidity(repository)
	getLatestMeasurementUseCase := use_case.NewGet_Latest(repository)

	// Instancia los controladores
	createController := controllers.NewCreate_TemperatureAndHumidity_C(createUseCase)
	getTemperatureByIDController := controllers.NewGetTemperatureByIDController(getTemperatureByIDUseCase)
	getAllController := controllers.NewGetAllTemperatureAndHumidityController(getAllUseCase)
	deleteController := controllers.NewDeleteTemperatureAndHumidityController(deleteUseCase)
	getAverageTemperatureController := controllers.NewGetAverageTemperatureController(getAverageTemperatureUseCase)
	getAverageHumidityController := controllers.NewGetAverageHumidityController(getAverageHumidityUseCase)
	getLatestMeasurementController := controllers.NewGetLatestMeasurementController(getLatestMeasurementUseCase)

	return createController,
		getTemperatureByIDController,
		getAllController,
		deleteController,
		getAverageTemperatureController,
		getAverageHumidityController,
		getLatestMeasurementController,
		repository,
		serviceNotification, // Añadido aquí
		nil
}
