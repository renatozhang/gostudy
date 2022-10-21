package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from consoile failed, err:%v\n", err)
		}

		data = strings.TrimSpace(data)

		_, err = conn.Write([]byte(data))
		if err != nil {
			fmt.Printf("write failed, err:%v\n", err)
			break
		}

	}
}
