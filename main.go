package main

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/loafoe/go-rabbitmq"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

//CrovelWorker
func CrovelWorker(exchange string) (rabbitmq.ConsumerHandlerFunc, error) {
	producer, err := rabbitmq.NewProducer(rabbitmq.Config{
		Exchange:     viper.GetString("dest_exchange"),
		ExchangeType: viper.GetString("dest_exchange_type"),
		Durable:      false,
	})
	if err != nil {
		fmt.Printf("Error creating producer: %v\n", err)
		return nil, err
	}

	return func(deliveries <-chan amqp.Delivery, done <-chan bool) {
		for {
			select {
			case msg := <-deliveries:
				_ = producer.Publish(exchange, msg.RoutingKey, amqp.Publishing{
					Headers:         msg.Headers,
					ContentEncoding: msg.ContentEncoding,
					CorrelationId:   msg.CorrelationId,
					Priority:        msg.Priority,
					ContentType:     msg.ContentType,
					Body:            msg.Body,
					// More fields here
				})
			}
		}
	}, nil
}

func main() {
	viper.SetEnvPrefix("crovel")
	viper.SetDefault("src_routing_key", "#")
	viper.SetDefault("src_exchange_type", "topic")
	viper.SetDefault("dest_exchange_type", "topic")
	viper.AutomaticEnv()

	id := uuid.New().String()
	worker, err := CrovelWorker(viper.GetString("dest_exchange"))
	if err != nil {
		fmt.Printf("Error creating worker: %v\n", err)
		return
	}

	consumer, err := rabbitmq.NewConsumer(rabbitmq.Config{
		RoutingKey:   viper.GetString("src_routing_key"),
		Exchange:     viper.GetString("src_exchange"),
		ExchangeType: viper.GetString("src_exchange_type"),
		Durable:      false,
		AutoDelete:   true,
		QueueName:    "crovel-" + id,
		CTag:         "crovel-not-shovel",
		HandlerFunc:  worker,
	})
	if err != nil {
		fmt.Printf("Error creating consumer: %v\n", err)
		return
	}

	err = consumer.Start()
	if err != nil {
		fmt.Printf("Error starting consumer: %v\n", err)
		return
	}

	select {}
}
