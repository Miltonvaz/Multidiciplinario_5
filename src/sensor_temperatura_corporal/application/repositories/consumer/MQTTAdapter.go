package consumer

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/application/use_case"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTAdapter struct {
	UseCase *use_case.Create_BodyTemperature
}

func NewMQTTAdapter(useCase *use_case.Create_BodyTemperature) *MQTTAdapter {
	return &MQTTAdapter{UseCase: useCase}
}

func (adapter *MQTTAdapter) HandleMessage(client mqtt.Client, msg mqtt.Message) {
	var sensor entities.BodyTemperature

	if err := json.Unmarshal(msg.Payload(), &sensor); err != nil {
		log.Printf("Error al deserializar los datos: %v\n", err)
		return
	}

	_, err := adapter.UseCase.Execute(sensor)
	if err != nil {
		log.Printf("Error al procesar los datos: %v\n", err)
		return
	}

	log.Printf("Datos procesados correctamente: %+v\n", sensor)
}
