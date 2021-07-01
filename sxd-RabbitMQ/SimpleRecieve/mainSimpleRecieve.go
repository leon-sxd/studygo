package main

import "sxd-RabbitMQ/RabbitMQ"

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "sxd")
	rabbitmq.ConsumeSimple()
}