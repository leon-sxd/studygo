package main

import (
	"fmt"
	"strconv"
	"sxd-RabbitMQ/RabbitMQ"
	"time"
)

func main()  {
	sxdone := RabbitMQ.NewRabbitMQTopic("exSxdTopic","sxd.topic.one")
	sxdtwo := RabbitMQ.NewRabbitMQTopic("exSxdTopic","sxd.topic.two")

	for i:=0; i<=100 ; i++  {
		sxdone.PublishTopic("hello sxd topic one ! "+strconv.Itoa(i))
		sxdtwo.PublishTopic("hello sxd topic two ! "+ strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
