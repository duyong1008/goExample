package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second) //延时1s
	}
}

func main() {

	//新建一个协程
	go newTask() //必须放到main之前

	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second) //延时1s
	}

}
