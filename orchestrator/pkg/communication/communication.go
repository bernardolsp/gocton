package communication

import (
	"context"
	"errors"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Communicator struct {
	Type             string
	ConnectionString string
}

type Messager struct {
	Channel *amqp.Channel
	Context context.Context
	Queue   string
}

func Initialize(c *Communicator) (Messager, error) {
	var messager Messager
	switch c.Type {
	case "rabbit":
		m, err := initialize_rabbitmq(c.ConnectionString, "test")
		if err != nil {
			log.Printf("Error initializing rabbitmq, %v", err)
			return messager, err
		}
		return m, err
	default:
		return messager, errors.New("unrecognized communicator. Cancelling execution")
	}
}

func (m *Messager) SendMessage(queue string, message []byte) error {
	err := m.Channel.PublishWithContext(m.Context,
		"",
		m.Queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		log.Println("Error publishing message to messager, ", err)
	}
	return nil
}

func (m *Messager) ReceiveMessage(queue string) ([]byte, error) {
	return nil, nil
}
