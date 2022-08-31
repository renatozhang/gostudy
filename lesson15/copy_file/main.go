package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open source file %s failed,err:%v\n", srcName, err)
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open dest file %s failed,err:%v\n", srcName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	_, err := CopyFile("target.txt", "main.go")
	if err != nil {
		fmt.Printf("copy file dailed, err:%v\n", err)
	}
	fmt.Println("copy done")
}
