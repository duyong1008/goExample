package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4)

	//接收多少写多少
	for {
		n, err := conn.Read(buf) //接收对方发过来的文件内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {

				fmt.Println("conn.Read err = ", err)
				return
			}
			return
		}
		f.Write(buf[:n]) //往文件写入内容
	}
}

func main() {

	//监听
	listenner, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	//阻塞等待用户连接
	conn, err1 := listenner.Accept()
	if err1 != nil {
		fmt.Println("listenner.Accept err = ", err)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)

	var n int
	n, err = conn.Read(buf) //读取对方发送的文件名
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}
	fileName := string(buf[:n])

	//回复“ok”
	conn.Write([]byte("ok"))

	//接收文件内容
	RecvFile(fileName, conn)
}
