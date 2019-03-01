package main

import (
	"fmt"
)

type Student struct {
	name string
	id   int
}

func main() {
	//定义切片为三个元素，类型为口接口
	i := make([]interface{}, 3)
	i[0] = 1                    //int
	i[1] = "hello go"           //string
	i[2] = Student{"mike", 666} //Student

	/*****************断言-if*****************/
	fmt.Printf("*****************断言-if*****************\n")
	//类型查询，类型断言
	//第一个返回下标，第二个返回下标对应的值，data分别是值
	for index, data := range i {
		//第一个返回值，第二个返回判断结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d]类型为int,内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d]类型为string,内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d]类型为Student,内容name = %s,id = %d\n", index, value.name, value.id)
		}
	}

	/*****************断言-switch*****************/
	fmt.Printf("*****************断言-switch*****************\n")
	for index, data := range i {
		switch value := data.(type) { //返回对应的值
		case int:
			fmt.Printf("x[%d]类型为int,内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d]类型为string,内容为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d]类型为Student,内容name = %s,id = %d\n", index, value.name, value.id)
		}
	}
}
