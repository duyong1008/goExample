package main

import "fmt"
import "github.com/jinzhu/configor"

func main() {
	fmt.Println("使用外部包测试：", configor.Config{})
}
