package adapter_m

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/application/repositories"
	"Multidiciplinario/src/sensor_temperatura_corporal/application/use_case"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type MQTTAdapter struct {
	UseCase *use_case.Create_BodyTemperature
	client  mqtt.Client
}

func NewMQTTAdapter(db domain.IBodyTemperature, serviceNotification *repositories.ServiceNotification) (*MQTTAdapter, error) {
	useCase := use_case.NewCreate_BodyTemperature(db, serviceNotification)

	adapter := &MQTTAdapter{UseCase: useCase}

	client, _, err := adapter.ConnectAndConsume()
	if err != nil {
		return nil, err
	}

	adapter.client = *client
	return adapter, nil
}

func loadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}
	return nil
}

func (adapter *MQTTAdapter) ConnectAndConsume() (*mqtt.Client, string, error) {
	if err := loadEnvVariables(); err != nil {
		return nil, "", err
	}

	brokerURL := os.Getenv("MQTT_BROKER_URL")
	clientID := os.Getenv("MQTT_CLIENT_ID")
	username := os.Getenv("MQTT_USERNAME")
	password := os.Getenv("MQTT_PASSWORD")
	topic := os.Getenv("MQTT_TOPIC")

	if brokerURL == "" || clientID == "" || username == "" || password == "" || topic == "" {
		return nil, "", fmt.Errorf("Missing environment variables for MQTT connection")
	}

	opts := mqtt.NewClientOptions().AddBroker(brokerURL).
		SetClientID(clientID).
		SetUsername(username).
		SetPassword(password).
		SetDefaultPublishHandler(adapter.HandleMessageAdapter)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, "", fmt.Errorf("Error connecting to MQTT broker: %v", token.Error())
	}
	token := client.Subscribe(topic, 0, adapter.HandleMessageAdapter)
	if token.Wait() && token.Error() != nil {
		return nil, "", fmt.Errorf("Error subscribing to topic: %v", token.Error())
	}

	log.Printf("Successfully connected to MQTT broker. Subscribed to topic: %s\n", topic)
	return &client, clientID, nil
}

func (adapter *MQTTAdapter) HandleMessageAdapter(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Message received on topic %s: %s\n", msg.Topic(), string(msg.Payload()))
	adapter.HandleMessage(msg)
}

func (adapter *MQTTAdapter) HandleMessage(msg mqtt.Message) {
	var sensor entities.BodyTemperature

	if err := json.Unmarshal(msg.Payload(), &sensor); err != nil {
		log.Printf("Error unmarshalling data: %v\n", err)
		return
	}
	createdSensor, err := adapter.UseCase.Execute(sensor)
	if err != nil {
		log.Printf("Error saving data: %v\n", err)
		return
	}

	log.Printf("Temperature and humidity data processed successfully: %+v\n", createdSensor)
}
