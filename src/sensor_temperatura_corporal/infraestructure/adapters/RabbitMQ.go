package adapters

import (
	"Multidiciplinario/src/sensor_temperatura_corporal/application/repositories"
	"Multidiciplinario/src/sensor_temperatura_corporal/domain/entities"
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

// NewRabbitMQAdapter initializes a new RabbitMQ connection and channel.
func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file:", err)
	}

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		return nil, fmt.Errorf("RABBITMQ_URL is not defined in the .env file")
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	if err := declareQueue(ch); err != nil {
		return nil, err
	}

	if err := ch.Confirm(false); err != nil {
		return nil, fmt.Errorf("failed to enable message confirmations: %w", err)
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

// declareQueue ensures the "sesnsor.temperatura" queue exists.
func declareQueue(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare(
		"sesnsor.temperatura",
		true,  // Durable queue
		false, // Auto-delete queue
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}
	return nil
}

// PublishEvent sends a BodyTemperature event to the RabbitMQ queue.
func (r *RabbitMQAdapter) PublishEvent(eventType string, data entities.BodyTemperature) error {
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal event data: %w", err)
	}

	ack, nack := r.ch.NotifyConfirm(make(chan uint64, 1), make(chan uint64, 1))

	err = r.ch.Publish(
		"",
		"sesnsor.temperatura",
		true,  // Persistent delivery mode
		false, // No immediate delivery
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	go handlePublishConfirmation(ack, nack)

	log.Printf("Event published: %s", eventType)
	return nil
}

// handlePublishConfirmation processes the message confirmation asynchronously.
func handlePublishConfirmation(ack, nack <-chan uint64) {
	select {
	case <-ack:
		log.Println("Message confirmed by RabbitMQ")
	case <-nack:
		log.Println("Message not confirmed by RabbitMQ")
	}
}

// Close gracefully shuts down the RabbitMQ connection and channel.
func (r *RabbitMQAdapter) Close() {
	if err := r.ch.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ channel: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ connection: %v", err)
	}
}
