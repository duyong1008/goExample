package main

import (
	"errors"
	"fmt"
)

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("被除数不能为零")
	} else {
		result = a / b
	}
	return
}

func main() {
	//方法一
	err1 := fmt.Errorf("%s", "this is mormol err")
	fmt.Println("err1=", err1)

	//方法二
	err2 := errors.New("this is normal err2")
	fmt.Println("err2=", err2)

	//方法三
	result, err := MyDiv(10, 2)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("rresult = ", result)
	}
}
