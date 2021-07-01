package main

import "sxd-RabbitMQ/RabbitMQ"

func main()  {
	sxdone := RabbitMQ.NewRabbitMQTopic("exSxdTopic","#")
	sxdone.ReceiveTopic()
}
