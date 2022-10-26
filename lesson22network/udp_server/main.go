package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	defer socket.Close()
	for {
		var date [1024]byte
		count, addr, err := socket.ReadFromUDP(date[:])
		if err != nil {
			fmt.Printf("read upd dailed, err:%v\n", err)
			continue
		}
		fmt.Printf("data:%s addr:%v count:%d\n", string(date[0:count]), addr, count)
		_, err = socket.WriteToUDP([]byte("hello client"), addr)
		if err != nil {
			fmt.Printf("read udp failed, err:%v\n", err)
			continue
		}
	}

}
