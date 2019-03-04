package main

import (
	"fmt"
	"time"
)

func nilCacheChannel() {
	//创建一个无缓存的channel
	ch := make(chan int)

	//len(ch)缓冲区剩余数据个数， cap（ch)缓冲区大小
	//fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))
	//新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程：i = %d\n", i)
			ch <- i //往chan写内容
			//fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))
		}
	}()

	//延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch //读管道中内容，没有内容前，阻塞
		fmt.Println("num = ", num)
	}
}

func cacheChannel() {
	//创建一个无缓存的channel
	ch := make(chan int, 3)

	//len(ch)缓冲区剩余数据个数， cap（ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))
	//新建协程
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i //往chan写内容
			fmt.Printf("子协程[%d]:len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))
		}
	}()

	//延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch //读管道中内容，没有内容前，阻塞
		fmt.Println("num = ", num)
	}
}

func closeChannel() {
	//创建一个无缓存Channel
	ch := make(chan int)

	//新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往通道写数据
		}
		//不需要再写数据时，关闭channel
		close(ch)
	}()

	//for死循环获取channel数据
	/*for {
		//如果ok为true,说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else { //管道关闭
			break
		}
	}*/

	//通过range获取channel数据
	for num := range ch {
		fmt.Println("num = ", num)
	}
}

//此通道只能写，不能读
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

//此通道只能读，不能写
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num = ", num)
	}
}

func aoneChannel() {
	//创建一个双向通道
	ch := make(chan int)

	//生产者，生产数字，写入channel
	//新开一个协程
	go producer(ch) //channel传参，引用传递

	//消费者，从channel读取内容，打印
	consumer(ch)
}

func main() {
	//无缓冲Channel
	//nilCacheChannel()

	//有缓冲Channel
	//cacheChannel()

	//关闭Channel
	//closeChannel()

	//单向Channel
	aoneChannel()

}
