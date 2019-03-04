package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
设备文件：
	屏幕（标准输出设备）	fmt.Println()	往标准输出设备写内容
	键盘（标准输入设备）	fmt.Scan()		从标准输入设备读取内容

磁盘文件，放在存储设备上的文件
	文本文件		以记事本打开，能看到内容（不是乱码）
	二进制文件	以记事本打开，能看到内容（是乱码）

*/

func deviceUse() {
	//are u ok?输出到屏幕
	//os.Stdout.Close()		//关闭后，无法输出
	fmt.Println("are u ok?") //住标准输出设备（屏幕）写内容
	//标准设备文件(os.Stdout)，默认已经打开，用户可以直接使用
	os.Stdout.WriteString("are u ok?\n") //和fmt.Println()功能一样

	//os.Stdin.Close()	//关闭后，无法输入
	var a int
	fmt.Println("请输入a: ")
	fmt.Scan(&a) //从标准输入设备中读取内容，放在a中
	fmt.Println("a = ", a)
}

func diskUse() {
	//写文件
	path := "./demo.txt"
	WriteFile(path)
	//读文件,按buf大小读取
	ReadFile(path)
	//读文件，每次读取一行
	ReadFileLine(path)
}

func WriteFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//使用完毕，需要关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i=%d\n", i)
		//fmt.Println("buf = ", buf)
		_, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		//fmt.Println("n = ", n)
	}
}

func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err =", err)
		return
	}

	//关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) //2k大小

	//n代表从文件读取内容的长度
	n, err1 := f.Read(buf)             //Read是按buf的大小来读取数据
	if err1 != nil && err1 != io.EOF { //io.EOF代表文件读取结束
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))

}

//每次读取一行
func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err =", err)
		return
	}

	//关闭文件
	defer f.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)
	for {

		//遇到'\n'结束读取,但是'\n'也读取进去了
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("bur = #%s#\n", string(buf))
	}

}

func copyFile() {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("Usage: xxx srcfile dstFile")
		return
	}
	srcFileName := list[1]
	dstFileNmae := list[2]
	if srcFileName == dstFileNmae {
		fmt.Println("源文件和目标文件不能相同")
		return
	}

	//只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	//新建目标文件
	dF, err2 := os.Create(dstFileNmae)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	//操作完毕，需要关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心处理，从源文件读取内容，写入目标文件，读多少写多少
	buf := make([]byte, 4*1024) //4k大小临时缓冲区
	for {
		n, err := sF.Read(buf) //从源文件读取内容
		if err != nil {
			if err == io.EOF { //文件读取完毕
				break
			}
			fmt.Println("err = ", err)
		}
		//住目标文件写，读多少写多少
		dF.Write(buf[:n])
	}
}

func main() {
	//设备文件的使用
	//deviceUse()

	//磁盘文件的使用	常用
	//diskUse()

	//案例，实现一个拷贝文件的功能
	copyFile()
}
