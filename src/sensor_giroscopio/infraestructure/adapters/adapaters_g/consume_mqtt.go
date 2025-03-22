package adapters_g

import (
	"Multidiciplinario/src/sensor_giroscopio/application/repositories"
	"Multidiciplinario/src/sensor_giroscopio/application/use_case"
	"Multidiciplinario/src/sensor_giroscopio/domain"
	entities "Multidiciplinario/src/sensor_giroscopio/domain/entinties"

	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type MQTTAdapter struct {
	UseCase *use_case.Create_GyroscopeSensor
	client  mqtt.Client
}

// NewMQTTAdapter initializes the MQTT adapter and connects to the broker
func NewMQTTAdapter(db domain.IGyroscopeSensor, serviceNotification *repositories.ServiceNotification) (*MQTTAdapter, error) {
	useCase := use_case.NewCreate_GyroscopeSensor(db, serviceNotification)
	adapter := &MQTTAdapter{UseCase: useCase}

	client, _, err := adapter.connectAndConsume()
	if err != nil {
		return nil, err
	}

	adapter.client = *client
	return adapter, nil
}

// loadEnvVariables loads the environment variables from the .env file
func loadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}
	return nil
}

// connectAndConsume connects to the MQTT broker and subscribes to the topic
func (adapter *MQTTAdapter) connectAndConsume() (*mqtt.Client, string, error) {
	if err := loadEnvVariables(); err != nil {
		return nil, "", err
	}

	brokerURL, clientID, username, password, topic := getMQTTConfig()
	if brokerURL == "" || clientID == "" || username == "" || password == "" || topic == "" {
		return nil, "", fmt.Errorf("Missing environment variables for MQTT connection")
	}

	opts := mqtt.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID(clientID).
		SetUsername(username).
		SetPassword(password).
		SetDefaultPublishHandler(adapter.handleMessageAdapter)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, "", fmt.Errorf("Error connecting to MQTT broker: %v", token.Error())
	}

	if token := client.Subscribe(topic, 0, adapter.handleMessageAdapter); token.Wait() && token.Error() != nil {
		return nil, "", fmt.Errorf("Error subscribing to topic: %v", token.Error())
	}

	log.Printf("Successfully connected to MQTT broker. Subscribed to topic: %s\n", topic)
	return &client, clientID, nil
}

// getMQTTConfig retrieves the MQTT configuration from environment variables
func getMQTTConfig() (string, string, string, string, string) {
	return os.Getenv("MQTT_BROKER_URL"),
		"GoSubscriber3", // clientID fixed for simplicity
		os.Getenv("MQTT_USERNAME"),
		os.Getenv("MQTT_PASSWORD"),
		"esp32.giroscopio" // topic fixed for simplicity
}

// handleMessageAdapter is the default handler for messages received on the subscribed topic
func (adapter *MQTTAdapter) handleMessageAdapter(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Message received on topic %s: %s\n", msg.Topic(), string(msg.Payload()))
	adapter.handleMessage(msg)
}

// handleMessage processes the incoming message and saves the data
func (adapter *MQTTAdapter) handleMessage(msg mqtt.Message) {
	var sensor entities.GyroscopeSensor

	if err := json.Unmarshal(msg.Payload(), &sensor); err != nil {
		log.Printf("Error unmarshalling data: %v\n", err)
		return
	}

	createdSensor, err := adapter.UseCase.Execute(sensor)
	if err != nil {
		log.Printf("Error saving data: %v\n", err)
		return
	}

	log.Printf("Gyroscope data processed successfully: %+v\n", createdSensor)
}
