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
		log.Println("RABBITMQ_URL no está definida en el archivo .env")
		return nil, fmt.Errorf("RABBITMQ_URL no está definida")
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Printf("Error conectando a RabbitMQ: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error abriendo canal: %v", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"sensor.ritmo_cardiaco",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Error declarando la cola: %v", err)
		return nil, err
	}

	err = ch.Confirm(false)
	if err != nil {
		log.Printf("Error habilitando confirmaciones de mensaje: %v", err)
		return nil, err
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

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

func (r *RabbitMQAdapter) Close() {
	if err := r.ch.Close(); err != nil {
		log.Printf("Error cerrando canal RabbitMQ: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Printf("Error cerrando conexión RabbitMQ: %v", err)
	}
}
