package adapters

import (
	"Multidiciplinario/src/sensor_luz/application/repositories"
	"Multidiciplinario/src/sensor_luz/domain/entities"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

var _ repositories.NotificationPort = (*RabbitMQAdapter)(nil)

// NewRabbitMQAdapter initializes a new RabbitMQ connection and channel, and declares a queue
func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	if err := loadEnv(); err != nil {
		return nil, err
	}

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		return nil, fmt.Errorf("RABBITMQ_URL is not defined in the .env file")
	}

	conn, err := connectToRabbitMQ(rabbitURL)
	if err != nil {
		return nil, err
	}

	ch, err := createChannel(conn)
	if err != nil {
		return nil, err
	}

	if err := declareQueue(ch); err != nil {
		return nil, err
	}

	if err := enableConfirmations(ch); err != nil {
		return nil, err
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

// loadEnv loads the environment variables from the .env file
func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load .env file, ensure it exists")
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

// connectToRabbitMQ establishes a connection to RabbitMQ
func connectToRabbitMQ(rabbitURL string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Printf("Error connecting to RabbitMQ: %v", err)
		return nil, err
	}
	return conn, nil
}

// createChannel creates a new RabbitMQ channel
func createChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error opening channel: %v", err)
		return nil, err
	}
	return ch, nil
}

// declareQueue declares the queue for message delivery
func declareQueue(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare(
		"sensor.luz",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Error declaring queue: %v", err)
		return err
	}
	return nil
}

// enableConfirmations enables message confirmations for the channel
func enableConfirmations(ch *amqp.Channel) error {
	if err := ch.Confirm(false); err != nil {
		log.Printf("Error enabling message confirmations: %v", err)
		return err
	}
	return nil
}

// PublishEvent publishes an event message to RabbitMQ
func (r *RabbitMQAdapter) PublishEvent(eventType string, data entities.LightSensorLDR) error {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling event to JSON: %v", err)
		return err
	}

	ack, nack := r.ch.NotifyConfirm(make(chan uint64, 1), make(chan uint64, 1))

	if err := r.ch.Publish(
		"",
		"sensor.luz",
		true,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		log.Printf("Error sending message to RabbitMQ: %v", err)
		return err
	}

	select {
	case <-ack:
		log.Println("Message confirmed by RabbitMQ")
	case <-nack:
		log.Println("Message not confirmed by RabbitMQ")
	}

	log.Println("Event published:", eventType)
	return nil
}

// Close closes the RabbitMQ channel and connection
func (r *RabbitMQAdapter) Close() {
	if err := closeChannel(r.ch); err != nil {
		log.Printf("Error closing RabbitMQ channel: %v", err)
	}
	if err := closeConnection(r.conn); err != nil {
		log.Printf("Error closing RabbitMQ connection: %v", err)
	}
}

// closeChannel closes the RabbitMQ channel
func closeChannel(ch *amqp.Channel) error {
	if err := ch.Close(); err != nil {
		return fmt.Errorf("error closing channel: %v", err)
	}
	return nil
}

// closeConnection closes the RabbitMQ connection
func closeConnection(conn *amqp.Connection) error {
	if err := conn.Close(); err != nil {
		return fmt.Errorf("error closing connection: %v", err)
	}
	return nil
}
