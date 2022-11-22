package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer(strings.Split("localhost:9092", ","), nil)
	if err != nil {
		fmt.Println("Failed to start consumenr: ", err)
		return
	}

	partitionList, err := consumer.Partitions("nginx_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList { // 遍历所有分区
		// 针对每一个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		// 异步从每个分区消费消息
		go func(pc1 sarama.PartitionConsumer) {
			for msg := range pc1.Messages() {
				fmt.Printf("partition:%d, offset:%d, key:%s, value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
			wg.Done()
		}(pc)

	}
	wg.Wait()
	consumer.Close()
}
