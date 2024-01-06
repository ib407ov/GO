package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	for {

		//Connect to RebbitMQ
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		//Error checking
		failOnError(err, "Failed to connect to RabbitMQ")
		//if err != nil {
		//	log.Fatalf("Failed to connect to RabbitMQ : %s", err)
		//}
		//Stop rebbitMQ after
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		// метод оголошує нову чергу або перевіряє існуючу чергу з заданими параметрами.
		q, err := ch.QueueDeclare(
			"text_queue", // Назва черги
			false,        // durable
			false,        // autoDelete
			false,        // exclusive
			false,        // noWait
			nil,          // arguments
		)
		failOnError(err, "Failed to declare a queue")

		fmt.Print("Введіть текст: ")
		var text string
		fmt.Scanln(&text)

		//Цей фрагмент коду відповідає за відправлення повідомлення до черги text_queue на сервері RabbitMQ через вже створений канал
		
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(text),
			})
		failOnError(err, "Failed to publish a message")

		log.Printf(" [x] Sent '%s'", text)
	}
}
