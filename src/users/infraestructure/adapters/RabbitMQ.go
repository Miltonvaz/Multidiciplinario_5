package adapters

import (
	"Multidiciplinario/src/users/application/repositories"
	"Multidiciplinario/src/users/domain/entities"
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
	if err := loadEnv(); err != nil {
		return nil, err
	}

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		return nil, fmt.Errorf("RABBITMQ_URL no está definido en el archivo .env")
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

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, asegúrate de que existe")
		return fmt.Errorf("error cargando .env: %v", err)
	}
	return nil
}

func connectToRabbitMQ(rabbitURL string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Printf("Error al conectar con RabbitMQ: %v", err)
		return nil, err
	}
	return conn, nil
}

func createChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error al abrir el canal: %v", err)
		return nil, err
	}
	return ch, nil
}

func declareQueue(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare(
		"esp32.datos",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Error al declarar la cola: %v", err)
		return err
	}
	return nil
}

func enableConfirmations(ch *amqp.Channel) error {
	if err := ch.Confirm(false); err != nil {
		log.Printf("Error al habilitar confirmaciones: %v", err)
		return err
	}
	return nil
}

func (r *RabbitMQAdapter) PublishEvent(eventType string, data entities.User) error {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error al convertir evento a JSON: %v", err)
		return err
	}

	ack, nack := r.ch.NotifyConfirm(make(chan uint64, 1), make(chan uint64, 1))

	if err := r.ch.Publish(
		"",
		"esp32.datos",
		true,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		log.Printf("Error al enviar mensaje a RabbitMQ: %v", err)
		return err
	}

	select {
	case <-ack:
		log.Println("Mensaje confirmado por RabbitMQ")
	case <-nack:
		log.Println("Mensaje NO confirmado por RabbitMQ")
	}

	log.Println("Evento publicado:", eventType)
	return nil
}

func (r *RabbitMQAdapter) Close() {
	if err := closeChannel(r.ch); err != nil {
		log.Printf("Error al cerrar canal de RabbitMQ: %v", err)
	}
	if err := closeConnection(r.conn); err != nil {
		log.Printf("Error al cerrar conexión de RabbitMQ: %v", err)
	}
}

func closeChannel(ch *amqp.Channel) error {
	if err := ch.Close(); err != nil {
		return fmt.Errorf("error al cerrar canal: %v", err)
	}
	return nil
}

func closeConnection(conn *amqp.Connection) error {
	if err := conn.Close(); err != nil {
		return fmt.Errorf("error al cerrar conexión: %v", err)
	}
	return nil
}
