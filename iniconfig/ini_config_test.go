package iniconfig

import (
	"fmt"
	"io/ioutil"
	"testing"
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

func TestIniConfig(t *testing.T) {
	fmt.Println("hello")
	data, err := ioutil.ReadFile("config.ini")
	if err != nil {
		t.Error("read fail failed")
	}

	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Errorf("unmarshal failed, err:%v\n", err)
	}
	t.Logf("unmarshal sucess, conf:%#v\n", conf)

	confData, err := Marshal(conf)
	if err != nil {
		t.Errorf("marshal failed,err:%v\n", err)
	}
	t.Logf("marshal sucess,conf:%s\n", string(confData))

	// MarshalFile(conf, "./test.conf")

}

func TestIniConfigFile(t *testing.T) {
	filename := "./test.conf"
	var conf Config
	conf.ServerConf.Ip = "localhost"
	conf.ServerConf.Port = 8888
	err := MarshalFile(conf, filename)
	if err != nil {
		t.Errorf("marshal failed, err:%v\n", err)
	}

	var conf2 Config
	err = UnMarshalFile(filename, &conf2)
	if err != nil {
		t.Errorf("unmarshal failed, err:%v\n", err)
	}
	t.Logf("unmarshal sucess,conf:%#v\n", conf2)
}
