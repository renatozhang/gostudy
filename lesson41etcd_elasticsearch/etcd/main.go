package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func testConnect() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"locahoshost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	defer cli.Close()
}

func testPut() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/conf", "simple_value")
	cancel()
	if err != nil {
		fmt.Println("put failed,err:", err)
		return
	}
	fmt.Println("put success")

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/conf")
	cancel()
	if err != nil {
		fmt.Println("get failed,err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("get %s: %s\n", ev.Key, ev.Value)
	}
}

func testDelete() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Delete(ctx, "/logagent/conf")
	cancel()
	if err != nil {
		fmt.Println("delete failed, err:", err)
		return
	}
	for _, v := range resp.PrevKvs {
		fmt.Printf("delete %s: %s\n", v.Key, v.Value)
	}
}

func testWatch() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect success")
	defer cli.Close()

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		_, err := cli.Put(ctx, "/logagent/conf", "test_watch")
		cancel()
		if err != nil {
			fmt.Println("put failed, err:", err)
			return
		}
	}()

	for {
		rch := cli.Watch(context.Background(), "/logagent/conf")
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("watch success, %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}

	}

}

func main() {
	// testConnect()
	// testPut()
	// testDelete()
	testWatch()
}
