package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	defer conn.Close()

	requestBuf := "GET /go HTTP/1.1\r\n" +
		"Host: 127.0.0.1:8000\r\n" +
		"Upgrade-Insecure-Requests: 1\r\n" +
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\n" +
		"User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.1.2 Safari/605.1.15\r\n" +
		"Accept-Language: zh-cn\r\n" +
		"Accept-Encoding: gzip, deflate\r\n" +
		"Connection: keep-alive\r\n" +
		"\r\n"

	//先发请求包，服务器才会回响应包
	conn.Write([]byte(requestBuf))

	//接收服务器回复的响应包
	buf := make([]byte, 1024*4)
	n, err1 := conn.Read(buf)
	if n == 0 {
		fmt.Println("conn.Read err1 =", err1)
		return
	}

	//打印响应报文
	fmt.Printf("#%v#", string(buf[:n]))

}
