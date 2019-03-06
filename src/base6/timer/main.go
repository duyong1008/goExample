package main

import (
	"fmt"
	"time"
)

func easyTimer() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())

	//2s后，往timer.C写数据，有数据后，就可以读取
	t := <-timer.C //channel没有数据前会阻塞

	fmt.Println("t = ", t)
}

func easyTimer2() {
	timer := time.NewTimer(1 * time.Second)

	for {
		<-timer.C
		fmt.Println("时间到")
	}
}

func easyTimer3() {
	//延时2s后打印一句话
	//1.通过timer实现
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("time.C时间到")

	//2.通过sleep实现
	time.Sleep(2 * time.Second)
	fmt.Println("timer.Sleep时间到")

	//3.time.After()实现
	time.After(2 * time.Second) //定时2s，阻塞2s，2s后产生一个事件，往channel写内容
	fmt.Println("time.After时间到")
}

func easyTimer4() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了，因为定时器的时间到")
	}()

	timer.Stop() //停止定时器

	for {
	}
}

func easyTimer5() {
	timer := time.NewTimer(3 * time.Second)
	timer.Reset(1 * time.Second) //重新设置为1s
	<-timer.C

	fmt.Println("时间到")

}

func easyTimer6() {
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for {
		<-ticker.C

		i++
		fmt.Println("i = ", i)

		if i == 5 {
			ticker.Stop()
			break
		}
	}
}

func fibonacci() {
	ch := make(chan int)    //数字通信
	quit := make(chan bool) //程序是否结束

	//消费者，从channel读取内容
	//新建协程
	go consumers(ch, quit)

	//生产者，产生数字，写入channel
	producers(ch, quit)
}

//ch只写，quit只读
func producers(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		//监听channel数据的流动
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}

//ch只读，quit只写
func consumers(ch <-chan int, quit chan<- bool) {
	for i := 0; i < 8; i++ {
		num := <-ch
		fmt.Println(num)
	}
	//可以停止
	quit <- true
}

func selectTimeOut() {
	ch := make(chan int)
	quit := make(chan bool)
	//新开一个协程
	go func() {
		for {
			select {
			//case num, isClose := <-ch:
			case num := <-ch:
				//if isClose != true {
				//	fmt.Println("Ch channel关闭")
				//	return
				//}
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
				close(quit)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	//close(ch)
	<-quit
	fmt.Println("程序结束")
}

func main() {
	//创建一个定时器，设置时间为2s，2s后，往time通道写内容（当前时间）
	//easyTimer()

	//验证time.NewTimer(),时间到了，只会响应一次
	//easyTimer2() //结果，只会响应一次

	//通过Timer实现延时功能
	//easyTimer3()

	//定时器停止
	//easyTimer4()

	//定时器重置
	//easyTimer5()

	//响应多次
	//easyTimer6()

	//通过select实现斐波那契数列
	//fibonacci()

	//通过select实现超时
	selectTimeOut()
}
