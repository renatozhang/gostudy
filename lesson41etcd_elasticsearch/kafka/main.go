package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition 随机分配partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("producer closr,err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)

}
