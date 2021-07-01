package main

import "sxd-RabbitMQ/RabbitMQ"

func main()  {
	sxdone := RabbitMQ.NewRabbitMQRouting("sxd","sxd_one")
	sxdone.ReceiveRouting()
}
