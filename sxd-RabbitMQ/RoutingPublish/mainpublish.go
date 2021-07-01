package main

import (
	"fmt"
	"strconv"
	"time"
	"sxd-RabbitMQ/RabbitMQ"
)

func main()  {
	sxdone := RabbitMQ.NewRabbitMQRouting("sxd","sxd_one")
	sxdtwo := RabbitMQ.NewRabbitMQRouting("sxd","sxd_two")

	for i:=0 ; i<=100 ; i++ {
		sxdone.PublishRouting("hello sxd one!"+strconv.Itoa(i))
		sxdtwo.PublishRouting("hello sxd two!"+strconv.Itoa(i+10))
		time.Sleep(1*time.Second)
		fmt.Println(i)
	}
}
