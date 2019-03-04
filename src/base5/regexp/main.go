package main

import (
	"fmt"
	"regexp"
)

func strMatch() [][]string {
	buf := "abc azc a7c aac 888 a9c tac"

	//1.解释规则,它会解析正则表达式，如果成功返回解释器
	//reg1 :=regexp.MustCompile(`a.c`)
	reg1 := regexp.MustCompile(`a[\d]c`)

	//2.根据规则提取关键信息, -1表式匹配所有
	test := reg1.FindAllStringSubmatch(buf, -1)
	return test
}

func floatMatch() [][]string {
	buf := "3.14 567 adfgg 1.23 7. 8.99 1sdfasdffg 6.666 7.8 888.8888"

	//解释正则表达式
	reg := regexp.MustCompile(`\d+\.\d+`)

	//提取关键信息
	result := reg.FindAllStringSubmatch(buf, -1)
	return result

}

func netMatch() [][]string {
	//反引号就是原生字符串
	buf := `
	<!DOCTYPE html>
	<html lang="zh-CN">
	<head>
		<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
		<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
		<meta charset="utf-8">
		<link rel="shortcut icon" href="/static/img/go.ico">
		<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
		<meta name="author" content="polaris <polaris@studygolang.com>">
		<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
		<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
	</head>
		<div>和爱好</div>
		<div>哈哈
			在不在
			不在
		</div>
		<div>测试</div>
		<div>你过来啊</div>
	<frameset cols="15,85">
		<frame src="/static/pkgdoc/i.html">
		<frame name="main" src="/static/pkgdoc/main.html" tppabs="main.html" >
		<noframes>
		</noframes>
	</frameset>
	</html>`

	//解释正则表达式
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)

	//提取关键信息
	result := reg.FindAllStringSubmatch(buf, -1)
	for _, text := range result {
		fmt.Println("text[0]=", text[0])   //带div
		fmt.Println("text[1] = ", text[1]) //内容
	}
	return result
}

func main() {
	//正则匹配字符串
	reuslt := strMatch()
	fmt.Println(reuslt)

	//正则匹配小数
	floatResult := floatMatch()
	fmt.Println(floatResult)

	//正则匹配web数据
	netResult := netMatch()
	fmt.Println("netResult = ", netResult)
}
