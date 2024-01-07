package broadcasterrepo

import (
	"context"
	"fmt"
	"time"

	"rpc-server/internal/core/domain/enum"
	"rpc-server/internal/port/outbound"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqBroadcasterRepository struct {
	ch           *amqp.Channel
	exchangeName string
}

func NewRabbitmqBroadcasterRepository(ch *amqp.Channel, exchangeName string) (outbound.BroadcasterRepository, error) {
	err := ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare an exchange: %w", err)
	}

	return &RabbitmqBroadcasterRepository{
		ch:           ch,
		exchangeName: exchangeName,
	}, nil
}

func (bc *RabbitmqBroadcasterRepository) Publish(color enum.Color) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return bc.ch.PublishWithContext(ctx,
		bc.exchangeName, // exchange
		"",              // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(color.String()),
		})
}

func (bc *RabbitmqBroadcasterRepository) Subscribe() (<-chan enum.Color, error) {
	q, err := bc.ch.QueueDeclare(
		"",    // name (automatically generated)
		false, // durable
		false, // delete when unused
		true,  // exclusive (auto-delete on close)
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %w", err)
	}

	err = bc.ch.QueueBind(
		q.Name,          // queue name
		"",              // routing key
		bc.exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind a queue: %w", err)
	}

	msgs, err := bc.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}

	// TODO: abstract this logic ⬇️ because there is exactly the same in `watermill.go`

	// Create a new channel for enum.Color messages
	colorChannel := make(chan enum.Color)

	// Start a goroutine to convert and forward messages to colorChannel
	go func() {
		defer close(colorChannel)
		for msg := range msgs {
			// Assuming you have a way to convert amqp.Delivery to enum.Color
			colorMsg, err := enum.ColorString(string(msg.Body))
			if err != nil {
				// Handle the error (e.g., log it) and continue
				continue
			}

			// Send the converted message to colorChannel
			colorChannel <- colorMsg
		}
	}()

	return colorChannel, nil
}
