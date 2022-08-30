package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./main.go.gz")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("gzip new reader failed, err:", err)
		return
	}
	var content []byte
	var buf [128]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file:", err)
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}
