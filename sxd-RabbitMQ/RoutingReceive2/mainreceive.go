package main

import "sxd-RabbitMQ/RabbitMQ"

func main()  {
	sxdtwo := RabbitMQ.NewRabbitMQRouting("sxd","sxd_two")
	sxdtwo.ReceiveRouting()
}
