package mq

import (
	"fmt"

	"github.com/Levap123/playstar-test/logging_service/logs"
	"github.com/streadway/amqp"
)

type Consumer struct {
	delivery <-chan amqp.Delivery
	logger   *logs.Logger
}

func NewConsumer(mqChan *amqp.Channel) (*Consumer, error) {
	messages, err := mqChan.Consume(
		"logs",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		delivery: messages,
	}, nil
}

func (c *Consumer) Consume() {
	for msg := range c.delivery {
		fmt.Println(msg.Body)
	}
}
