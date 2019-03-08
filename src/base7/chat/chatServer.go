package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	Msg  chan string //用于发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

//保存在线用户	cliAddr ====> Client
var onlineMap map[string]Client

var message = make(chan string)

//处理用户链接
func HandleConn(conn net.Conn) {
	//
	defer conn.Close()

	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体,默认，用户名和网络地址一样
	cli := Client{
		make(chan string),
		cliAddr,
		cliAddr,
	}

	//把结构体添加到map
	onlineMap[cliAddr] = cli

	//新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient(cli, conn)

	//广播某个人在线
	message <- MakeMsg(cli, "login")

	//提示，我是谁
	cli.Msg <- MakeMsg(cli, "I am here")

	//对方是否主动退出
	isQuit := make(chan bool)

	//对方是否有数据发送
	hasData := make(chan bool)

	//新建一个协程，接收用户发送过来的数据
	go func() {
		buf := make([]byte, 2048)
		for {

			n, err := conn.Read(buf)
			if n == 0 { //对方断开或者出问题
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := buf[:n] //通过windows nc测试，多一个换行需要n-1

			//清除字符中的\n\r
			conv := strings.Trim(string(msg), "\n\r")

			if len(conv) == 3 && "who" == conv {
				//遍历map，给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg := tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(conv) >= 8 && "rename" == conv[:6] {
				//rename|dy		重命名
				name := strings.Split(conv, "|")
				if len(name) < 2 {
					conn.Write([]byte("rename failure\n"))
					continue
				}
				cli.Name = name[1]
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))

			} else if len(conv) == 4 && "exit" == conv {
				isQuit <- true
			} else {
				//转发此内容
				message <- MakeMsg(cli, string(msg))
			}

			//代表有数据
			hasData <- true

		}
		fmt.Println("dddrerer")
		return
	}()

	for {
		//通过select检测channel的流动
		select {
		case <-isQuit:
			//当前用户从map移除
			delete(onlineMap, cliAddr)

			//广播谁下线了
			message <- MakeMsg(cli, "login out")
			return
		case <-hasData:

		case <-time.After(60 * time.Second): //60s后超时
			delete(onlineMap, cliAddr)
			message <- MakeMsg(cli, "time out leave out") //广播谁下线了
			return
		}
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	//给当前客户端发送信息
	for msg := range cli.Msg {
		conn.Write([]byte(msg + "\n"))
	}
}

//新开一个协程，转发消息，只要有消息来了，遍历map,给map每个成员发送消息
func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)
	for {
		msg := <-message //没有消息前，这里会阻塞

		//遍历map,给map每个成员发送消息
		for _, cli := range onlineMap {
			cli.Msg <- msg
		}
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen = ", err)
		return
	}

	defer listener.Close()

	//新开一个协程，转发消息，只要有消息来了，遍历map,给map每个成员发送消息
	go Manager()

	//主协和，循环阻塞等待用户链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		go HandleConn(conn) //处理用户链接
	}

}
