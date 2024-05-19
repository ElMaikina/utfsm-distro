package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the queues
	queues := []string{"async_queue1", "async_queue2", "sync_queue"}
	for _, queue := range queues {
		_, err = ch.QueueDeclare(
			queue, // name
			false, // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			log.Fatalf("Failed to declare a queue: %v", err)
		}
	}

	// Listen to async queues asynchronously
	for _, queue := range queues[:2] {
		go listenAsync(ch, queue)
	}

	// Listen to the sync queue synchronously
	listenSync(ch, "sync_queue")

	// Block the main thread to keep the program running
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("Shutting down")
}

// listenAsync listens to a RabbitMQ queue asynchronously
func listenAsync(ch *amqp.Channel, queue string) {
	msgs, err := ch.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received a message from %s: %s", queue, msg.Body)
		}
	}()
}

// listenSync listens to a RabbitMQ queue synchronously
func listenSync(ch *amqp.Channel, queue string) {
	msgs, err := ch.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for msg := range msgs {
		log.Printf("Received a message from %s: %s", queue, msg.Body)
	}
}
