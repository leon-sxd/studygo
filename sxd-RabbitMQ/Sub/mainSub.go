package main

import "sxd-RabbitMQ/RabbitMQ"

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	rabbitmq.ReceiveSub()
}