package main

import (
	"fmt"
	"strconv"
	"sxd-RabbitMQ/RabbitMQ"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")

	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("订阅模式生成第 " + strconv.Itoa(i) + " 条数据")
		fmt.Println("fmt: 订阅模式生成第 " + strconv.Itoa(i) + " 条数据")
		time.Sleep(1 * time.Second)
	}
}
