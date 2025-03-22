package adapters

import (
	"Multidiciplinario/src/sensor_ritmo_cardiaco/application/repositories"
	"Multidiciplinario/src/sensor_ritmo_cardiaco/domain/entities"
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

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, asegurate de que existe")
	}

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		return nil, fmt.Errorf("RABBITMQ_URL no está definida en el archivo .env")
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, fmt.Errorf("Error conectando a RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Error abriendo canal: %v", err)
	}

	if err := declareQueue(ch); err != nil {
		return nil, err
	}

	if err := enableConfirmations(ch); err != nil {
		return nil, err
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

// declareQueue encapsulates the logic to declare the RabbitMQ queue
func declareQueue(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare(
		"sensor.ritmo_cardiaco", // queue name
		true,                    // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	if err != nil {
		log.Printf("Error declarando la cola: %v", err)
	}
	return err
}

// enableConfirmations enables message confirmations on the channel
func enableConfirmations(ch *amqp.Channel) error {
	if err := ch.Confirm(false); err != nil {
		log.Printf("Error habilitando confirmaciones de mensaje: %v", err)
		return err
	}
	return nil
}

// PublishEvent publishes a HeartRate event to RabbitMQ
func (r *RabbitMQAdapter) PublishEvent(eventType string, data entities.HeartRate) error {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error convirtiendo evento a JSON: %v", err)
		return err
	}

	ack, nack := r.ch.NotifyConfirm(make(chan uint64, 1), make(chan uint64, 1))

	err = r.ch.Publish(
		"",
		"sensor.ritmo_cardiaco",
		true,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("Error enviando mensaje a RabbitMQ: %v", err)
		return err
	}

	select {
	case <-ack:
		log.Println("Mensaje confirmado por RabbitMQ")
	case <-nack:
		log.Println("Mensaje no fue confirmado")
	}

	log.Println("Evento publicado:", eventType)
	return nil
}

// Close closes the RabbitMQ connection and channel
func (r *RabbitMQAdapter) Close() {
	if err := r.ch.Close(); err != nil {
		log.Printf("Error cerrando canal RabbitMQ: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Printf("Error cerrando conexión RabbitMQ: %v", err)
	}
}
