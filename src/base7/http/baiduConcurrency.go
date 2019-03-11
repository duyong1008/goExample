package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet1(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	//读取网页body内容
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { //读取结束或出问题
			if err == io.EOF {
				fmt.Println("数据爬取完成")
				break
			}
			fmt.Println("resp.Body.Read err = ", err)
			break
		}

		result += string(buf[:n])
	}
	return
}

//爬取一个网页
func SpiderPage(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Printf("正在爬第%d页网页%s\n", i, url)

	//2.爬内容
	result, err := HttpGet1(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}

	//把内容写入到文件
	fileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create err1 = ", err1)
		return
	}

	//写内容
	f.WriteString(result)

	f.Close()

	page <- i

}

func DoWork1(start int, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	//创建管道
	page := make(chan int)
	//1.明确目标
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页（ >= 1 ) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= %d) :", start)
	fmt.Scan(&end)

	DoWork1(start, end)
}
