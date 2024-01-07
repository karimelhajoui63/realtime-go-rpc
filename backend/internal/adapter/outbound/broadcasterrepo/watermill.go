package broadcasterrepo

import (
	"context"
	"fmt"
	"log"
	"time"

	"rpc-server/internal/core/domain/enum"
	"rpc-server/internal/port/outbound"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

type WattermillBroadcasterRepository struct {
	subscriber   *amqp.Subscriber
	publisher    *amqp.Publisher
	exchangeName string
}

func NewWatermillBroadcasterRepository(cfg *amqp.Config, exchangeName string) (outbound.BroadcasterRepository, error) {
	subscriber, err := amqp.NewSubscriber(
		// This config is based on this example: https://www.rabbitmq.com/tutorials/tutorial-two-go.html
		// It works as a simple queue.
		//
		// If you want to implement a Pub/Sub style service instead, check
		// https://watermill.io/pubsubs/amqp/#amqp-consumer-groups
		*cfg,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a subscriber: %w", err)
	}

	publisher, err := amqp.NewPublisher(*cfg, watermill.NewStdLogger(false, false))
	if err != nil {
		return nil, fmt.Errorf("failed to declare a publisher: %w", err)
	}

	return &WattermillBroadcasterRepository{
		subscriber:   subscriber,
		publisher:    publisher,
		exchangeName: exchangeName,
	}, nil
}

func (bc *WattermillBroadcasterRepository) Publish(color enum.Color) error {
	msg := message.NewMessage(watermill.NewUUID(), []byte(color.String()))
	return bc.publisher.Publish(bc.exchangeName, msg)
}

func (bc *WattermillBroadcasterRepository) Subscribe() (<-chan enum.Color, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msgs, err := bc.subscriber.Subscribe(ctx, bc.exchangeName)
	if err != nil {
		return nil, fmt.Errorf("failed to register a subscriber: %w", err)
	}

	// Create a new channel for enum.Color messages
	colorChannel := make(chan enum.Color)

	// Start a goroutine to convert and forward messages to colorChannel
	go func() {
		defer close(colorChannel)
		for msg := range msgs {
			// Assuming you have a way to convert message.Message to enum.Color
			colorMsg, err := enum.ColorString(string(msg.Payload))
			if err != nil {
				log.Println("Could not convert this to enum.Color:", string(msg.Payload))
				continue
			}

			// Send the converted message to colorChannel
			colorChannel <- colorMsg
		}
	}()

	return colorChannel, nil
}
