package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

type Consumer struct{}

func (c Consumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive:", msg.NSQDAddress, "message:", string(msg.Body))
	return nil

}

func main() {
	//nsq的地址
	nsqAddress := "127.0.0.1:4161"
	err := initConsumer("order_queue", "first", nsqAddress)
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}

//初始化消费者
func initConsumer(topic, channel, address string) error {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second     //设置服务发现的轮询时间
	c, err := nsq.NewConsumer(topic, channel, config) //新建一个消费者
	if err != nil {
		return err
	}
	consumer := &Consumer{}
	c.AddHandler(consumer)
	if err := c.ConnectToNSQLookupd(address); err != nil {
		return err
	}
	return nil
}
