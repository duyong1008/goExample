package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"` //二次编码		,如果使用`json:"-"`则不打印此字段
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok,string"` //使用,string可用转换值为字符串类型
	Price    float64  `json:"price"`
}

//只解析需要的字段
type IT2 struct {
	Subjects []string `json:"subjects"`
}

func codingJson() {
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
	buf, err := json.MarshalIndent(s, "", "	")
	if err != nil {
		return
	}
	fmt.Println("struct to json = ", string(buf))
}

func decodingJson() {
	jsonBuf := `
	{
    "Company":"itcast",
    "Subjects":[
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "Isok":"true",
    "Price":666.666
	}`

	var tmp1 IT                                    //定义一个结构体变量
	err1 := json.Unmarshal([]byte(jsonBuf), &tmp1) //第二个参数要地址传递
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	fmt.Println("json conv struct tmp1 = ", tmp1)

	var tmp2 IT2
	err2 := json.Unmarshal([]byte(jsonBuf), &tmp2) //第二个参数要地址传递
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}
	fmt.Println("json conv struct tmp2指定字段 = ", tmp2)
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
	fmt.Println("map conv json =  ", string(result))
}

func jsonFormap() {
	jsonBuf := `
	{
    "Company":"itcast",
    "Subjects":[
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "Isok":"true",
    "Price":666.666
	}`
	//创建一个map
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Printf("m = %+v\n", m)

	//类型断言, 值，它是value类型
	for key, value := range m {
		//fmt.Printf("%v ======= %v\n", key, value)
		switch data := value.(type) {
		case string:
			fmt.Printf("map[%s]的值类型为string,value = %s\n", key, data)
		case bool:
			fmt.Printf("map[%s]的值类型为bool,value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64,value = %f\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]interface{},value = %v\n", key, data)
		}

	}
}

func main() {
	//通过结构体生成编码json
	codingJson()

	//通过json解码生成结构体
	//1.通过json解析到结构体，虽然前期写结构体很麻烦，但后面要获取值一目了然
	decodingJson()

	//通过map生成json
	mapForJson()

	//通过json解码生成到map
	//1.通过json解析到map很方便，但要获取值需要类型断言，所以很麻烦
	jsonFormap()

}
