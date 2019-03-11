package main

import (
	"fmt"
	"net"
	"net/http"
)

//GET请求 明文
//#GET / HTTP/1.1
//Host: 127.0.0.1:8000
//Upgrade-Insecure-Requests: 1
//Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
//User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.1.2 Safari/605.1.15
//Accept-Language: zh-cn
//Accept-Encoding: gzip, deflate
//Connection: keep-alive
//

func getRequest() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listener.Close()
	//阻塞等待用户的连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err1 = ", err1)
		return
	}

	defer conn.Close()

	//接收客户端的数据
	buf := make([]byte, 1024*4)
	n, err2 := conn.Read(buf)
	if n == 0 { //对方断开或出问题
		fmt.Println("conn.Read err2 = ", err2)
		return
	}

	fmt.Printf("#%v#", string(buf[:n]))
}

//服务器编写的业务逻辑处理程序
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func getResponse() {
	http.HandleFunc("/go", myHandler)

	//在指定的地址进行监听，开启一个HTTP
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func getRequestBaidu() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status = ", resp.Status)
	fmt.Println("StatusCode = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)
	//fmt.Println("Body = ", resp.Body)
	buf := make([]byte, 4*1024)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read err = ", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("tmp = ", tmp)

}

func main() {
	//GET请求包
	//getRequest()

	//GET响应包
	//getResponse()

	//GET请求baidu
	getRequestBaidu()

}
