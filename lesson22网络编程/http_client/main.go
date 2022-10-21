package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:443")
	if err != nil {
		fmt.Println("Errror dialing", err.Error())
		return
	}

	defer conn.Close()
	data := "GET / HTTP/1.1\r\n"
	data += "HOST: www.baidu.com\r\n"
	data += "Connection: close\r\n"
	data += "\r\n\r\n"

	_, err = io.WriteString(conn, data)
	if err != nil {
		fmt.Println("write string failed, err:", err)
		return
	}

	buf := make([]byte, 4096)
	for {
		fmt.Println("start read")
		count, err := conn.Read(buf)
		if err != nil || count == 0 {
			break
		}
		fmt.Println(string(buf[:count]))
	}
}
