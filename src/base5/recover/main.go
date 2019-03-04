package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaa")
}

func testb(i int) {
	//设置recover
	//recover()	//可以打印panic的错误信息
	defer func() {
		err := recover() //产生了panic异常
		if err != nil {
			fmt.Println(err)
		}

	}()
	var a [10]int
	a[i] = 111 //当x为20时候，导致数组下标越界，产生一个panic,导致程序崩溃
	fmt.Println("bbbbbbbbbbbb")
}

func testc() {
	fmt.Println("cccccccccccc")
}

func main() {
	testa()
	testb(12)
	testc()

}
