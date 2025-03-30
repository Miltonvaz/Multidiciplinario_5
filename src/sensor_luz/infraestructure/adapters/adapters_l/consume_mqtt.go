package adapters_l

import (
	"Multidiciplinario/src/sensor_luz/application/repositories"
	"Multidiciplinario/src/sensor_luz/application/use_case"
	"Multidiciplinario/src/sensor_luz/domain"
	"Multidiciplinario/src/sensor_luz/domain/entities"
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type MQTTAdapter struct {
	UseCase *use_case.Create_LightLDR
	client  mqtt.Client
}

func NewMQTTAdapter(db domain.ILightSensor, serviceNotification *repositories.ServiceNotification) (*MQTTAdapter, error) {
	useCase := use_case.NewCreate_LightLDR(db, serviceNotification)

	adapter := &MQTTAdapter{UseCase: useCase}

	client, err := adapter.ConnectAndConsume()
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

func (adapter *MQTTAdapter) ConnectAndConsume() (*mqtt.Client, error) {
	if err := loadEnvVariables(); err != nil {
		return nil, err
	}

	brokerURL, clientID, username, password, topic := getMQTTConfig()

	if brokerURL == "" || clientID == "" || username == "" || password == "" || topic == "" {
		return nil, fmt.Errorf("Missing environment variables for MQTT connection")
	}

	client, err := adapter.connectClient(brokerURL, clientID, username, password)
	if err != nil {
		return nil, err
	}

	err = adapter.subscribeToTopic(client, topic)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully connected to MQTT broker. Subscribed to topic: %s\n", topic)
	return &client, nil
}

func getMQTTConfig() (string, string, string, string, string) {
	return os.Getenv("MQTT_BROKER_URL"), "GoSubscriber2", os.Getenv("MQTT_USERNAME"), os.Getenv("MQTT_PASSWORD"), "esp32.luz"
}

func (adapter *MQTTAdapter) connectClient(brokerURL, clientID, username, password string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID(clientID).
		SetUsername(username).
		SetPassword(password).
		SetDefaultPublishHandler(adapter.HandleMessageAdapter)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("Error connecting to MQTT broker: %v", token.Error())
	}
	return client, nil
}

func (adapter *MQTTAdapter) subscribeToTopic(client mqtt.Client, topic string) error {
	token := client.Subscribe(topic, 0, adapter.HandleMessageAdapter)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("Error subscribing to topic: %v", token.Error())
	}
	return nil
}

func (adapter *MQTTAdapter) HandleMessageAdapter(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Message received on topic %s: %s\n", msg.Topic(), string(msg.Payload()))
	adapter.HandleMessage(msg)
}

func (adapter *MQTTAdapter) HandleMessage(msg mqtt.Message) {
	var sensor entities.LightSensorLDR

	if err := json.Unmarshal(msg.Payload(), &sensor); err != nil {
		log.Printf("Error unmarshalling data: %v\n", err)
		return
	}
	createdSensor, err := adapter.UseCase.Execute(sensor)
	if err != nil {
		log.Printf("Error saving data: %v\n", err)
		return
	}

	log.Printf("Light sensor data processed successfully: %+v\n", createdSensor)
}
