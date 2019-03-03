package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"` //二次编码
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok,string"` //使用,string可用转换值为字符串类型
	Price    float64  `json:"price"`
}

//只解析需要的字段
type IT2 struct {
	Subjects []string `json:"subjects"`
}

func codingJson() (buf []byte, err error) {
	//定义一个结构体变量
	s := IT{
		"itcast",
		[]string{"Go", "C++", "Python", "Test"},
		true,
		666.666,
	}

	//编码，根据内容生成json文本
	//buf,err :=json.Marshal(s)
	//格式化编码
	buf, err = json.MarshalIndent(s, "", "	")
	if err != nil {
		return
	}
	return
}

func decodingJson() (stt interface{}, err error) {
	jsonBuf := `
	{
    "Company":"itcast",
    "Subjects":[
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "Isok":true,
    "Price":666.666
	}`

	var tmp IT                                  //定义一个结构体变量
	err = json.Unmarshal([]byte(jsonBuf), &tmp) //第二个参数要地址传递
	if err != nil {
		return
	}

	var tmp2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &tmp2) //第二个参数要地址传递
	if err != nil {
		return
	}
	stt = tmp2
	return
}

func mapForJson() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 666.666

	//编码成json
	//result, err := json.Marshal(m)
	//编码格式化json
	result, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		return
	}
	fmt.Println(string(result))
}

func main() {
	//通过结构体生成编码json
	result, err := codingJson()
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("result = ", string(result))

	//通过json解码生成结构体
	result1, err1 := decodingJson()
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	fmt.Printf("result1 = %+v\n", result1.(IT2))

	//通过map生成json
	mapForJson()
}
