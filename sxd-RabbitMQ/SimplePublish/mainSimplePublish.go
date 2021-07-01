package main

import (
	"fmt"
	"strconv"
	"sxd-RabbitMQ/RabbitMQ"
	"time"
)

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "sxd")
	// simple模式
	//rabbitmq.PublishSimple("hello sxd 222!")
	//fmt.Println("发送成功")
	// work模式
	for i:=0;i<=100;i++{
		rabbitmq.PublishSimple("hello sxd !" +strconv.Itoa(i))
		time.Sleep(1*time.Second)
		fmt.Println(i)
	}
}
