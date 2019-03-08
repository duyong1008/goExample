package main

import (
	"fmt"
	"net"
	"strings"
)

func serverTCP() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		//处理用户请求,新建一个协程
		go HandleConn(conn)

	}
}

func HandleConn(conn net.Conn) {
	//函数调用完毕，自动关闭conn
	defer conn.Close()

	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "conncet sucessful")

	//读取用户数据
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		//打印接收的数据长度
		//fmt.Println(len(buf[:n]))

		//for index, data := range buf[:n] {
		//	fmt.Printf("%d:%c\n", index, data)
		//}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))

		if "exit" == string(buf[:n-1]) {
			fmt.Println(addr, "exit")
			return
		}

		//把数据转换为大写，在给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {

	//接收多个用户
	serverTCP()

}
