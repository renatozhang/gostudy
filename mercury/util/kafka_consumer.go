package util

import (
	"strings"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/renatozhang/gostudy/mercury/logger"
)

var wg sync.WaitGroup

func InitKafkaConsumer(addr, topic string, consume func(message *sarama.ConsumerMessage)) (err error) {
	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		logger.Error("Failed to start consumenr: ", err)
		return
	}

	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		logger.Error("Failed to get the list of partitions: ", err)
		return
	}

	logger.Debug("partition list:%#v", partitionList)
	for partition := range partitionList { // 遍历所有分区
		// 针对每一个分区创建一个对应的分区消费者
		pc, errRet := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if errRet != nil {
			err = errRet
			logger.Error("Failed to start consumer for partition %d: %v\n", partition, errRet)
			return
		}
		wg.Add(1)
		// 异步从每个分区消费消息
		go func(pc1 sarama.PartitionConsumer) {
			for msg := range pc1.Messages() {
				logger.Debug("partition:%d, offset:%d, key:%s, value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				consume(msg)
			}
			wg.Done()
		}(pc)
	}
	return
}
