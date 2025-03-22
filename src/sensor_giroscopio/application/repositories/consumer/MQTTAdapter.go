package consumer

import (
	"Multidiciplinario/src/sensor_giroscopio/application/use_case"
	entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTAdapter struct {
	UseCase *use_case.Create_GyroscopeSensor
}

func NewMQTTAdapter(useCase *use_case.Create_GyroscopeSensor) *MQTTAdapter {
	return &MQTTAdapter{UseCase: useCase}
}

func (adapter *MQTTAdapter) HandleMessage(client mqtt.Client, msg mqtt.Message) {
	var sensor entities.GyroscopeSensor

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
