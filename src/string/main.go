package main

import (
	"fmt"
	"strconv"
	"strings"
)

func strProcess() {

	//Contains 包含
	//“hellogo”中是否包含“hello”,包含返回true,不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))

	//Joins 合并
	//以“@”做连接符，返回一个新的字符串
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "@")
	fmt.Println("buf = ", buf)

	//Index
	//查找"hello"在“abcdhello”的索引位置，返回索引位置，如果不存在则返回-1
	fmt.Println(strings.Index("abcdhello", "hello"))

	//Repeat
	//重复多少次
	buf = strings.Repeat("go", 3)
	fmt.Println("buf = ", buf) //"gogogo"

	//Split
	//以指定的分隔符拆分字符串,返回切片
	buf = "hello@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2 = ", s2)

	//Trim
	//去掉两头的字符串空格
	buf = strings.Trim("     are u ok?        ", " ")
	fmt.Printf("buf = #%s#\n", buf)

	//Fields
	//去掉空格
	//只处理以空格的字符串做拆分，返回切片
	s3 := strings.Fields("     are u ok?        ")
	fmt.Println(s3)
}

func strConversion() {
	//定义一个字节切片
	slice := make([]byte, 0, 1024)
	//追加布尔类型字符串返回字节类型
	slice = strconv.AppendBool(slice, true)
	//第二个数为要追加的数，第3个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcgohello")
	fmt.Println(string(slice))

	//其它类型转换为字符串
	var str string

	str = strconv.FormatBool(false)
	//'f' 指打印格式，以小数方式， -1指精度（小数点位数）64指以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	//整型转换字符串，常用
	str = strconv.Itoa(6666)
	fmt.Println("str = ", str)

	//字符串转换其它类型
	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}

	//把字符串转换为整型	常用
	a, _ := strconv.Atoi("567")
	fmt.Println("a = ", a)
}

func main() {
	//字符串处理
	//strProcess()

	//字符串转换
	strConversion()
}
