package main

import (
	"fmt"
	"interfc/inherit"
)

type Humaner interface {
	sayhi()
}

type Student struct {
	name string
	id   int
}

type Teacher struct {
	addr  string
	group string
}

type MyStr string

//Student实现了此方法
func (this *Student) sayhi() {
	fmt.Printf("Student[%s,%d] sayhi\n", this.name, this.id)
}

//Teacher实现了此方法
func (this *Teacher) sayhi() {
	fmt.Printf("Teacher[%s,%s] sayhi\n", this.addr, this.group)
}

//MyStr实现了此方法
func (this *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi\n", *this)
}

//定义一个普通函数，函数的参数为接口类型
//只有一个函数，可以有不同表现，那么这个叫多态
func WhoSayHi(i Humaner) {
	i.sayhi()
}

func main() {
	//实例化三个变量
	s := &Student{"mike", 666}
	t := &Teacher{"wh", "golang"}
	var str MyStr = "hello world!"

	//调用同一函数，不同表现，多态，多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	//第一个返回下标，第二个返回下标对应的值
	for _, i := range x {
		i.sayhi()
	}

	/*****************继承*****************/

	//定义一个接口类型的变量
	var i inherit.Personer
	sd := &inherit.Student{"ddr333", 1111}
	i = sd
	i.Sayhi() //继承过来的方法
	i.Sing("学习歌")

}

/*
func main() {

	//定义接口类型变量
	var i Humaner

	//只要实现了此接口方法的类型，那么这个类型的变量就可以给i赋值
	s := &Student{"mike", 666}
	i = s
	s.sayhi()

	t := &Teacher{"wh", "go"}
	i = t
	t.sayhi()

	var str MyStr = "hello mike"
	i = &str
	i.sayhi()

}
*/
