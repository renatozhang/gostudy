package main

import (
	"fmt"

	"github.com/renatozhang/gostudy/iniconfig"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MysqlConfig struct {
	Username string  `ini:"username"`
	Password string  `ini:"password"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	TimeOut  float32 `ini:"timeout"`
}

func main() {
	filename := "/tmp/config.ini"
	var conf Config
	err := iniconfig.UnMarshalFile(filename, &conf)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}
	fmt.Printf("conf:%#v\n", conf)

	iniconfig.MarshalFile(conf, "/tmp/config1.ini")
}
