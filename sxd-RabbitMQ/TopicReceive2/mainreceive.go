package main

import "sxd-RabbitMQ/RabbitMQ"

func main()  {
	sxdtwo := RabbitMQ.NewRabbitMQTopic("exSxdTopic","sxd.*.two")
	sxdtwo.ReceiveTopic()
}
