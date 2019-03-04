package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccccccccccc")
	//return 	//终止此函数
	runtime.Goexit() //终止所在的协程
	fmt.Println("ddddddddddddd")
}

func main() {
	go func() {
		fmt.Println("aaaaaaaaaaa")
		//调用了别的函数
		test()
		fmt.Println("bbbbbbbbbbbbbbbbb")
	}()
	//特地写一个死循环，目的不让主协程结束
	for {

	}
}
