package util

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/renatozhang/gostudy/mercury/logger"
)

var (
	producer sarama.SyncProducer
)

func InitKafka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition 随机分配partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	producer, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logger.Error("producer close,err:", err)
		return
	}
	return
}

func SendKafka(topic string, value interface{}) (err error) {
	data, err := json.Marshal(value)
	if err != nil {
		return
	}
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	err = InitKafka("localhost:9092")
	if err != nil {
		fmt.Println("producer closr,err:", err)
		return
	}
	defer producer.Close()
	// 发送消息
	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		logger.Error("send message failed,", err)
		return
	}
	logger.Debug("pid:%v offset:%v, data:%#v\n", pid, offset, data)
	return
}
