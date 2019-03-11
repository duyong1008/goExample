package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HttpGet2(url string) (result string, err error) {
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

func SpiderOneJoy(url string) (title, content string, err error) {
	//开始爬取页面内容
	result, err1 := HttpGet2(url)
	if err1 != nil {
		err = err1
		return
	}

	//取关键信息
	//取标题	规则 “<h1>化缘															</h1>”
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile title err")
		return
	}
	//取标题内容
	tmpTitle := re1.FindAllStringSubmatch(result, 1) //最后一个参数为1，只过滤第一个
	for _, data := range tmpTitle {
		title = data[1]
		title = strings.Replace(title, "\t", "", -1)
		title = strings.Trim(title, " ")
		break
	}

	//取内容	规则 “<div class="content-txt pt10"> 内容 <a id="prev" href="”
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href=`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile content err")
		return
	}
	//取标题内容
	tmpContent := re2.FindAllStringSubmatch(result, 1) //最后一个参数为1，只过滤第一个
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\r", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "&nbsp;", "", -1)
		break
	}
	return

}

//爬取一个网页
func SpiderPage1(i int, page chan int) {
	url := "https://www.pengfue.com/index_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬第%d页网页%s\n", i, url)

	//开始爬取页面内容
	result, err := HttpGet2(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}

	//fmt.Println("r =", result)
	//取关键字<h1 class="dp-b"><a href=
	//正则解释表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}

	//获取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1)

	//定义一个切片
	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)
	//取网址
	//第一个返回下标，第二个返回内容
	for _, data := range joyUrls {
		//fmt.Println("Urls = ", data[1])
		//开始爬取每一个段子
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println("SpiderOneJoy err =", err)
			continue
		}
		//fmt.Printf("title = #%v#\n", title)
		//fmt.Printf("content = #%v#\n", content)
		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)

	}
	//把内容写入到文件
	StoreJoyToFile(i, fileTitle, fileContent)
	page <- i //写内容，写Num

}

//把内容写入到文件
func StoreJoyToFile(i int, fileTitle, fileContent []string) {

	fileName := strconv.Itoa(i) + ".txt"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create err1 = ", err1)
		return
	}

	defer f.Close()

	//写内容
	n := len(fileTitle)
	for i := 0; i < n; i++ {
		//写标题
		f.WriteString(fileTitle[i] + "\n")
		//写内容
		f.WriteString(fileContent[i] + "\n")
		//分割内容
		f.WriteString("\n==================================\n")
	}
}

func DoWork2(start int, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	//创建管道
	page := make(chan int)

	for i := start; i <= end; i++ {
		//定义一个函数，爬主页面
		go SpiderPage1(i, page)
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

	DoWork2(start, end)
}
