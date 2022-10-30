package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
	_, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	// r, err := redis.String(c.Do("lpop", "book_list")) // 后进先出
	r, err := redis.String(c.Do("rpop", "book_list")) // 先进后出

	if err != nil {
		fmt.Println("get book_list failed,", err)
		return
	}
	fmt.Println(r)
}
