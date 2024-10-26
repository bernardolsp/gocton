package communication

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func initialize_rabbitmq(c string, q string) (Messager, error) {
	conn, err := amqp.Dial(c)
	if err != nil {
		log.Fatalf("Error dialing rabbitmq, %v", err)
		return Messager{nil, nil, ""}, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error accessing channel, %v", err)
		return Messager{nil, nil, ""}, err
	}

	_, err = ch.QueueDeclare(
		q,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error declaring queue, %v", err)
		return Messager{nil, nil, ""}, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	return Messager{ch, ctx, q}, nil
}
