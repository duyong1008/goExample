package main

import (
	"fmt"
	"time"
)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello")
}

func person2() {
	Printer("world")
}

func main() {
	//新建2个协程，代表2个人，2个人同时使用打印
	go person1()
	go person2()

	//特地不让主协程结束，死循环
	for {

	}
}
